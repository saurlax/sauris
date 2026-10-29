// Command sauris-server is the Sauris control-plane entry point.
package main

import (
	"os"

	"github.com/saurlax/sauris/internal/command"
)

func main() {
	os.Exit(command.Run("sauris-server", os.Args[1:], os.Stdout, os.Stderr))
}
