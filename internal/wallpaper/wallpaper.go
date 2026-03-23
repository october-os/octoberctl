package wallpaper

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/october-os/octoberctl/internal/utils"
)

// Absolute path to the wallpapers folder.
var octoberWallDir string = fmt.Sprintf("%s/.config/october-config/wallpapers", os.Getenv("HOME"))

// ArgParser Parses the arguments given to the wallpaper flag.
func ArgParser(listWalls bool, addWall, removeWall, showWall string) error {
	if listWalls {
		if err := list(); err != nil {
			return err
		}
	}

	if addWall != "" {
		if err := add(addWall); err != nil {
			return err
		}
	}

	if removeWall != "" {
		absPath := fmt.Sprintf("%s/%s", octoberWallDir, removeWall)
		if err := utils.Remove(absPath); err != nil {
			return err
		}
	}

	if showWall != "" {
		absPath := fmt.Sprintf("%s/%s", octoberWallDir, showWall)
		if err := utils.Show(absPath); err != nil {
			return err
		}
	}

	return nil
}

// Adds the given file to the wallpapers folder.
func add(src string) error {
	if !utils.FileExist(src) {
		return fmt.Errorf("file %s does not exist", src)
	}

	fileIn, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fileIn.Close()

	fileInStat, _ := fileIn.Stat()

	dest := fmt.Sprintf("%s/%s", octoberWallDir, fileInStat.Name())
	if utils.FileExist(dest) {
		return fmt.Errorf("file %s already exist in the wallpapers folder", fileInStat.Name())
	}

	fileOut, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer fileOut.Close()

	_, err = io.Copy(fileOut, fileIn)
	return err
}

// List all the wallpapers in the wallpapers folder.
func list() error {
	wallpapers, err := os.ReadDir(octoberWallDir)
	if err != nil {
		return err
	}

	currentWallpaper, err := getCurrentWallpaper()
	if err != nil {
		return err
	}

	for i, entry := range wallpapers {
		if entry.Name() != ".gitkeep" {
			fmt.Printf("%d:\t%s", i, entry.Name())
			if entry.Name() == currentWallpaper {
				fmt.Printf("\t[current]")
			}
			fmt.Printf("\n")
		}
	}

	return nil
}

// Gets the current wallpaper inside
// /tmp/october-config/lastwallpaper.
func getCurrentWallpaper() (string, error) {
	lastWallPath := "/tmp/october-config/lastwallpaper"
	file, err := os.Open(lastWallPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var currentWallpaper string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentWallpaper = scanner.Text()
	}

	return currentWallpaper, nil
}
