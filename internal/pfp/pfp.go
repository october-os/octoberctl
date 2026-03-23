package pfp

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/october-os/octoberctl/internal/utils"
)

// Absolute path to pfp.
var octoberPfp string = fmt.Sprintf("%s/.config/october-config/profile_picture.jpg", os.Getenv("HOME"))

func ArgParser(peek, remove bool, set string) error {
	if peek {
		if err := utils.Show(octoberPfp); err != nil {
			return err
		}
	}

	if remove {
		if err := removePfp(); err != nil {
			return err
		}
	}

	if set != "" {
		if err := setPfp(set); err != nil {
			return err
		}
	}

	return nil
}

func removePfp() error {
	if err := utils.Remove(octoberPfp); err != nil {
		return fmt.Errorf("couldn't remove profile picture")
	}

	return nil
}

func setPfp(pfp string) error {
	if !utils.FileExist(pfp) || !strings.Contains(pfp, ".jpg") {
		return fmt.Errorf("file %s does not exist or isn't a jpg", pfp)
	}

	fileIn, err := os.Open(pfp)
	if err != nil {
		return err
	}
	defer fileIn.Close()

	fileOut, err := os.Create(octoberPfp)
	if err != nil {
		return err
	}
	defer fileOut.Close()

	_, err = io.Copy(fileOut, fileIn)
	return err
}
