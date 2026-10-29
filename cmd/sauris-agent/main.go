// Command sauris-agent is the Sauris host-agent entry point.
package main

import (
	"os"

	"github.com/saurlax/sauris/internal/command"
)

func main() {
	os.Exit(command.Run("sauris-agent", os.Args[1:], os.Stdout, os.Stderr))
}
