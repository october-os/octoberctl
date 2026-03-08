package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/october-os/octoberctl/internal/update"
	"github.com/october-os/octoberctl/internal/wallpaper"
)

var updateFlagSet *flag.FlagSet
var forcePtr bool

var wallpaperFlagSet *flag.FlagSet

var listWalls bool
var addWall string
var removeWall string
var showWall string

// Inits all the flags, args of the program and help messages.
func initFlags() {
	updateFlagSet = flag.NewFlagSet("update", flag.ExitOnError)
	updateFlagSet.BoolVar(&forcePtr, "f", false, "Force the update even if local changes are found")

	wallpaperFlagSet = flag.NewFlagSet("wallpaper", flag.ExitOnError)
	wallpaperFlagSet.BoolVar(&listWalls, "l", false, "List all wallpapers")
	wallpaperFlagSet.StringVar(&addWall, "a", "", "Add a wallpaper")
	wallpaperFlagSet.StringVar(&removeWall, "r", "", "Remove a wallpaper")
	wallpaperFlagSet.StringVar(&showWall, "s", "", "Show a wallpaper with kitten icat")

	flag.Usage = func() {
		mainHelpMessageHeader()
		flag.PrintDefaults()
	}

	updateFlagSet.Usage = func() {
		updateHelpMessageHeader()
		updateFlagSet.PrintDefaults()
	}

	wallpaperFlagSet.Usage = func() {
		wallpaperHelpMessageHeader()
		wallpaperFlagSet.PrintDefaults()
	}
}

// Parses the command that was used and its flags, then executes it.
func main() {
	initFlags()
	flag.Parse()

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(2)
	}

	subArgs := os.Args[2:] // Remove program name and command
	switch os.Args[1] {
	case "update":
		updateFlagSet.Parse(subArgs)
		if err := update.Update(forcePtr); err != nil {
			fmt.Printf("failed to update: %s\n", err.Error())
			os.Exit(1)
		}
	case "wallpaper":
		wallpaperFlagSet.Parse(subArgs)
		if err := wallpaper.WallpaperArgParser(listWalls, addWall, removeWall, showWall); err != nil {
			fmt.Println(err.Error())
			os.Exit(3)
		}
	default:
		fmt.Printf("error: unknown command '%s'\n", os.Args[1])
	}
}

// Prints the main help message header.
func mainHelpMessageHeader() {
	fmt.Print("usage: octoberctl <command> [<args>]\n\n")
	fmt.Println("commands:")
	fmt.Println("\tupdate\t\tUpdate the October Linux configuration")
	fmt.Println("\twallpaper\tManage wallpapers in the wallpapers folder")
}

// Prints the help message for the 'update' command.
func updateHelpMessageHeader() {
	fmt.Print("usage: octoberctl update [-f]\n\n")
	fmt.Println("args:")
}

// Prints the help message for 'wallpaper' command.
func wallpaperHelpMessageHeader() {
	fmt.Print("usage: octoberctl wallpaper [-l] [-a <path>] [-r <file name>] [-s <file name>]\n\n")
	fmt.Println("args:")
}
