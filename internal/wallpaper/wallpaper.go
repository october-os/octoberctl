package wallpaper

import (
	"fmt"
	"io"
	"os"
)

var octoberWallDir string = fmt.Sprintf("%s/.config/october-config/wallpapers", os.Getenv("HOME"))

func WallpaperArgParser(listWalls bool, addWall, removeWall string) error {
	if listWalls {
		if err := listWallpapers(); err != nil {
			return err
		}
	}

	if addWall != "" {
		if err := addWallpaper(addWall); err != nil {
			return err
		}
	}

	return nil
}

func addWallpaper(src string) error {
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
		return fmt.Errorf("File %s already exist in wallpapers folder", fileInStat.Name())
	}

	fileOut, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer fileOut.Close()

	_, err = io.Copy(fileOut, fileIn)
	return err
}

func removeWallpaper(wall string) {

}

func listWallpapers() error {
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
