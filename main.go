package main

import (
	_ "embed"

	"github.com/Dtx-validator/dtx-node/command/root"
	"github.com/Dtx-validator/dtx-node/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
