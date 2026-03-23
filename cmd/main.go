package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/october-os/octoberctl/internal/pfp"
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

var pfpFlagSet *flag.FlagSet
var setPfp string
var removePfp bool
var peekPfp bool

// Inits all the flags, args of the program and help messages.
func initFlags() {
	updateFlagSet = flag.NewFlagSet("update", flag.ExitOnError)
	updateFlagSet.BoolVar(&forcePtr, "f", false, "Force the update even if local changes are found")

	wallpaperFlagSet = flag.NewFlagSet("wallpaper", flag.ExitOnError)
	wallpaperFlagSet.BoolVar(&listWalls, "l", false, "List all wallpapers")
	wallpaperFlagSet.StringVar(&addWall, "a", "", "Add a wallpaper")
	wallpaperFlagSet.StringVar(&removeWall, "r", "", "Remove a wallpaper")
	wallpaperFlagSet.StringVar(&showWall, "s", "", "Show a wallpaper with kitten icat")

	pfpFlagSet = flag.NewFlagSet("pfp", flag.ExitOnError)
	pfpFlagSet.StringVar(&setPfp, "s", "", "Set the profile picture")
	pfpFlagSet.BoolVar(&removePfp, "r", false, "Remove the profile picture")
	pfpFlagSet.BoolVar(&peekPfp, "p", false, "See the profile picture")

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

	pfpFlagSet.Usage = func() {
		pfpHelpMessageHeader()
		pfpFlagSet.PrintDefaults()
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
		handleWallpaper(subArgs)
	case "pfp":
		handlePfp(subArgs)
	default:
		fmt.Printf("error: unknown command '%s'\n", os.Args[1])
	}
}

// Handles the operations for the wallpaper flag set.
func handleWallpaper(subArgs []string) {
	if len(subArgs) == 0 {
		wallpaperFlagSet.Usage()
		os.Exit(2)
	}

	wallpaperFlagSet.Parse(subArgs)
	if err := wallpaper.ArgParser(listWalls, addWall, removeWall, showWall); err != nil {
		fmt.Println(err.Error())
		os.Exit(3)
	}
}

// Handles the operations for the pfp flag set.
func handlePfp(subArgs []string) {
	if len(subArgs) == 0 {
		pfpFlagSet.Usage()
		os.Exit(2)
	}

	pfpFlagSet.Parse(subArgs)
	if err := pfp.ArgParser(peekPfp, removePfp, setPfp); err != nil {
		fmt.Println(err.Error())
		os.Exit(4)
	}
}

// Prints the main help message header.
func mainHelpMessageHeader() {
	fmt.Print("usage: octoberctl <command> [<args>]\n\n")
	fmt.Println("commands:")
	fmt.Println("\tupdate\t\tUpdate the October Linux configuration")
	fmt.Println("\twallpaper\tManage wallpapers")
	fmt.Println("\tpfp\t\tManage the profile picture")
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

func pfpHelpMessageHeader() {
	fmt.Print("usage: octoberctl pfp [-s <path>] [-r] [-p]\n\n")
	fmt.Println("args:")
}
