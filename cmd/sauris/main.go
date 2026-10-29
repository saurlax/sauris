// Command sauris is the Sauris CLI entry point.
package main

import (
	"os"

	"github.com/saurlax/sauris/internal/command"
)

func main() {
	os.Exit(command.Run("sauris", os.Args[1:], os.Stdout, os.Stderr))
}
