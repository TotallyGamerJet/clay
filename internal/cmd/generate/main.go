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
	m, err := parseAST(ast)
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

	// Build the Wasm module.
	wasm := filepath.Join(tmp, "clay.wasm")
	if _, err := output(clang, "--target=wasm32", "-std=c17", "-ffreestanding", "-nostdlib",
		"-O2", "-g0",
		"-mbulk-memory", "-msign-ext", "-mnontrapping-fptoint", "-mmutable-globals",
		"-I.", "-Ilibc",
		"-Wl,--no-entry", "-Wl,--strip-debug",
		"-o", wasm,
		"internal/wasm/src/clay.c", glue, "libc/malloc_sbrk.c", "libc/libc.c",
	); err != nil {
		log.Fatal(err)
	}

	// Translate it to Go.
	if _, err := output("go", "tool", "wasm2go", "-pkg", "wasm", "-unsafe", "-o", "internal/wasm/clay.go", wasm); err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile("clay.go", goSrc, 0o644); err != nil {
		log.Fatal(err)
	}
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
