/*
Copyright © 2025 nanvenomous mrgarelli@gmail.com
*/
package main

import (
	"embed"

	"github.com/nanvenomous/ssrStarter/cmd"
)

//go:embed build/*
var buildFS embed.FS

func main() {
	cmd.Execute(buildFS)
}
