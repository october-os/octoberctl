package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/october-os/octoberctl/internal/update"
)

// Parses the command that was used and its flags, then executes it.
func main() {
	updateFlagSet := flag.NewFlagSet("update", flag.ExitOnError)
	forcePtr := updateFlagSet.Bool("f", false, "Force the update even if changes were made")

	flag.Parse()

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(2)
	}

	subArgs := os.Args[2:] // Remove program name and command
	switch os.Args[1] {
	case "update":
		updateFlagSet.Parse(subArgs)
		if err := update.Update(*forcePtr); err != nil {
			fmt.Printf("failed to update: %s", err.Error())
		}
	default:
		fmt.Printf("error: unknown command '%s'\n", os.Args[1])
	}
}
