package main

import (
	"os"

	"github.com/eopenio/workwx/cmd/workwxctl/commands"
)

func main() {
	app := commands.InitApp()
	app.Run(os.Args)
}
