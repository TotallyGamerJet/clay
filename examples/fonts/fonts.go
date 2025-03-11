package fonts

import (
	_ "embed"
)

var (
	//go:embed Roboto-Regular.ttf
	RobotoRegularTTF []byte

	//go:embed RobotoMono-Medium.ttf
	RobotoMonoMediumTTF []byte
)
