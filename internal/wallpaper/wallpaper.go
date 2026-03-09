package wallpaper

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/blacktop/go-termimg"
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
		if err := remove(removeWall); err != nil {
			return err
		}
	}

	if showWall != "" {
		if err := show(showWall); err != nil {
			return err
		}
	}

	return nil
}

// Adds the given file to the wallpapers folder.
func add(src string) error {
	if !fileExist(src) {
		return fmt.Errorf("file %s does not exist", src)
	}

	fileIn, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fileIn.Close()

	fileInStat, _ := fileIn.Stat()

	dest := fmt.Sprintf("%s/%s", octoberWallDir, fileInStat.Name())
	if fileExist(dest) {
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

// Removes the given wallpaper from the wallpapers folder.
func remove(wall string) error {
	absPath := fmt.Sprintf("%s/%s", octoberWallDir, wall)
	if !fileExist(absPath) {
		return fmt.Errorf("wallpaper %s doesn't exist in the wallpapers folder", wall)
	}

	return os.Remove(absPath)
}

// Shows the given wallpaper in the terminal.
func show(wall string) error {
	absPath := fmt.Sprintf("%s/%s", octoberWallDir, wall)
	if !fileExist(absPath) {
		return fmt.Errorf("wallpaper %s doesn't exist in the wallpapers folder", wall)
	}

	image, err := termimg.Open(absPath)
	if err != nil {
		return err
	}

	err = image.WidthPixels(640).HeightPixels(360).Scale(termimg.ScaleAuto).Print()
	if err != nil {
		return err
	}

	fmt.Printf("\n")
	return nil
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

// Returns if the given file exist or not.
func fileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return !os.IsNotExist(err)
	}

	return true
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
