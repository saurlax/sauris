// Package command provides the bootstrap command behavior shared by Sauris binaries.
// It will not own deployment, networking, or other platform business logic.
package command

import (
	"fmt"
	"io"

	"github.com/saurlax/sauris/internal/buildinfo"
)

// Run handles the scaffold's help and version commands and returns an exit code.
func Run(name string, args []string, stdout, stderr io.Writer) int {
	if len(args) == 1 {
		switch args[0] {
		case "version", "--version":
			fmt.Fprintf(stdout, "%s %s\n", name, buildinfo.Version)
			return 0
		case "help", "--help", "-h":
			usage(name, stdout)
			return 0
		}
	}
	if len(args) == 0 {
		usage(name, stdout)
		return 0
	}
	fmt.Fprintln(stderr, "unsupported command: this scaffold only provides help and version")
	usage(name, stderr)
	return 2
}

func usage(name string, out io.Writer) {
	fmt.Fprintf(out, "Usage: %s [version|--version|help|--help]\n", name)
	fmt.Fprintln(out, "Sauris is in development. Platform services are not implemented yet.")
}
