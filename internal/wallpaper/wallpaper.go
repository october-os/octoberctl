package wallpaper

import (
	"fmt"
	"io"
	"os"

	"github.com/blacktop/go-termimg"
)

var octoberWallDir string = fmt.Sprintf("%s/.config/october-config/wallpapers", os.Getenv("HOME"))

func WallpaperArgParser(listWalls bool, addWall, removeWall, showWall string) error {
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

func add(src string) error {
	if !fileExist(src) {
		return fmt.Errorf("File %s does not exist", src)
	}

	fileIn, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fileIn.Close()

	fileInStat, _ := fileIn.Stat()

	dest := fmt.Sprintf("%s/%s", octoberWallDir, fileInStat.Name())
	if fileExist(dest) {
		return fmt.Errorf("File %s already exist in the wallpapers folder", fileInStat.Name())
	}

	fileOut, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer fileOut.Close()

	_, err = io.Copy(fileOut, fileIn)
	return err
}

func remove(wall string) error {
	absPath := fmt.Sprintf("%s/%s", octoberWallDir, wall)
	if !fileExist(absPath) {
		return fmt.Errorf("Wallpaper %s doesn't exist in the wallpapers folder", wall)
	}

	return os.Remove(absPath)
}

func show(wall string) error {
	absPath := fmt.Sprintf("%s/%s", octoberWallDir, wall)
	if !fileExist(absPath) {
		return fmt.Errorf("Wallpaper %s doesn't exist in the wallpapers folder", wall)
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

func list() error {
	wallpapers, err := os.ReadDir(octoberWallDir)
	if err != nil {
		return err
	}

	for i, entry := range wallpapers {
		if entry.Name() != ".gitkeep" {
			fmt.Printf("%d:\t%s\n", i, entry.Name())
		}
	}

	return nil
}

func fileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return !os.IsNotExist(err)
	}

	return true
}
