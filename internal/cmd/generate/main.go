// SPDX-License-Identifier: Zlib

// Command generate builds clay.h into Go.
//
// It compiles clay.h to WebAssembly with clang and translates the result to Go with wasm2go.
// Clang's AST of clay.h is used to generate the Go types that mirror clay's types,
// along with C and Go glue code that passes them in and out of the WebAssembly memory.
//
// It must be run from the root of the repository, which is what `go generate` does.
//
// Clang is looked up in the following order:
//   - the CLANG environment variable
//   - tools/wasi-sdk/bin/clang
//   - $WASI_SDK_PATH/bin/clang
//   - clang in PATH
package main

import (
	"bytes"
	"errors"
	"flag"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("generate: ")
	keep := flag.Bool("keep", false, "keep the intermediate files and print their location")
	flag.Parse()

	clang, err := findClang()
	if err != nil {
		log.Fatal(err)
	}

	tmp, err := os.MkdirTemp("", "clay-generate")
	if err != nil {
		log.Fatal(err)
	}
	if *keep {
		log.Printf("intermediate files are in %s", tmp)
	} else {
		defer os.RemoveAll(tmp)
	}

	// Dump the AST of the public part of clay.h.
	header := filepath.Join(tmp, "header.c")
	if err := os.WriteFile(header, []byte("#include \"clay.h\"\n"), 0o644); err != nil {
		log.Fatal(err)
	}
	ast, err := output(clang, "--target=wasm32", "-std=c17", "-ffreestanding", "-fsyntax-only",
		"-Xclang", "-ast-dump=json", "-I.", header)
	if err != nil {
		log.Fatal(err)
	}
	src, err := os.ReadFile("clay.h")
	if err != nil {
		log.Fatal(err)
	}
	m, err := parseAST(src, ast)
	if err != nil {
		log.Fatal(err)
	}
	goSrc, cSrc, err := generate(m)
	if err != nil {
		if goSrc != nil {
			os.WriteFile(filepath.Join(tmp, "clay.go"), goSrc, 0o644)
		}
		log.Fatal(err)
	}

	glue := filepath.Join(tmp, "glue.c")
	if err := os.WriteFile(glue, cSrc, 0o644); err != nil {
		log.Fatal(err)
	}

	// Extract the C standard library from the version of wasm2go in go.mod.
	libc := filepath.Join(tmp, "libc")
	if _, err := output("go", "tool", "libc-gen", "-c-out", libc); err != nil {
		log.Fatal(err)
	}

	// Build the Wasm module.
	wasm := filepath.Join(tmp, "clay.wasm")
	if _, err := output(clang, "--target=wasm32", "-std=c17", "-ffreestanding", "-nostdlib",
		"-O2", "-g0",
		"-mbulk-memory", "-msign-ext", "-mnontrapping-fptoint", "-mmutable-globals",
		"-I.", "-I"+libc,
		"-Wl,--no-entry", "-Wl,--strip-debug",
		"-o", wasm,
		"internal/wasm/src/clay.c", glue, filepath.Join(libc, "malloc_sbrk.c"), filepath.Join(libc, "libc.c"),
	); err != nil {
		log.Fatal(err)
	}

	// Translate it to Go.
	const translated = "internal/wasm/clay.go"
	if _, err := output("go", "tool", "wasm2go", "-pkg", "wasm", "-unsafe", "-o", translated, wasm); err != nil {
		log.Fatal(err)
	}
	if err := license(translated); err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile("clay.go", goSrc, 0o644); err != nil {
		log.Fatal(err)
	}
}

// license adds the license of the generated file, which is clay's, below the line
// that marks it as generated.
func license(path string) error {
	src, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	const notice = "// SPDX-License-Identifier: Zlib\n"
	if bytes.Contains(src, []byte(notice)) {
		return nil
	}
	generated, rest, _ := bytes.Cut(src, []byte("\n"))
	out := make([]byte, 0, len(src)+len(notice)+1)
	out = append(out, generated...)
	out = append(out, '\n')
	out = append(out, notice...)
	out = append(out, rest...)
	return os.WriteFile(path, out, 0o644)
}

func findClang() (string, error) {
	if c := os.Getenv("CLANG"); c != "" {
		return c, nil
	}
	candidates := []string{filepath.Join("tools", "wasi-sdk", "bin", "clang")}
	if sdk := os.Getenv("WASI_SDK_PATH"); sdk != "" {
		candidates = append(candidates, filepath.Join(sdk, "bin", "clang"))
	}
	for _, c := range candidates {
		if _, err := os.Stat(c); err == nil {
			return c, nil
		}
	}
	if c, err := exec.LookPath("clang"); err == nil {
		return c, nil
	}
	return "", errors.New("clang not found: install wasi-sdk into tools/wasi-sdk or set CLANG")
}

func output(name string, args ...string) ([]byte, error) {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command(name, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		os.Stderr.Write(stderr.Bytes())
		return nil, err
	}
	os.Stderr.Write(stderr.Bytes())
	return stdout.Bytes(), nil
}
