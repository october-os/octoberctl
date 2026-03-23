package utils

import (
	"fmt"
	"os"

	"github.com/blacktop/go-termimg"
)

// Returns if the given file exist or not.
func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return !os.IsNotExist(err)
	}

	return true
}

// Shows the given image in the terminal.
func Show(img string) error {
	if !FileExist(img) {
		return fmt.Errorf("image %s doesn't exist in the configuration", img)
	}

	image, err := termimg.Open(img)
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

// Removes the given file from the configuration.
func Remove(file string) error {
	if !FileExist(file) {
		return fmt.Errorf("%s doesn't exist in the configuration", file)
	}

	return os.Remove(file)
}
