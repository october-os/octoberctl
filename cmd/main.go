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
	forcePtr := updateFlagSet.Bool("f", false, "Force the update even if local changes are found")

	flag.Usage = func() {
		mainHelpMessageHeader()
		flag.PrintDefaults()
	}

	updateFlagSet.Usage = func() {
		updateHelpMessageHeader()
		updateFlagSet.PrintDefaults()
	}

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
			fmt.Printf("failed to update: %s\n", err.Error())
			os.Exit(1)
		}
	default:
		fmt.Printf("error: unknown command '%s'\n", os.Args[1])
	}
}

// Prints the main help message header.
func mainHelpMessageHeader() {
	fmt.Print("usage: octoberctl <command> [<args>]\n\n")
	fmt.Println("commands:")
	fmt.Println("\tupdate\t  Update the October Linux configuration")
}

// Prints the help message for the 'update' command.
func updateHelpMessageHeader() {
	fmt.Print("usage: octoberctl update [-f]\n\n")
	fmt.Println("args:")
}
