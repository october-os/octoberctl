package update

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/go-git/go-git/v5"
)

// Updates the October Linux configuration.
func Update(force bool) error {
	fmt.Println("Updating October configuration...")

	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	octoberConfigDir := fmt.Sprintf("%s/october-config", userConfigDir)

	repository, err := git.PlainOpen(octoberConfigDir)
	if err != nil {
		return err
	}

	workTree, err := repository.Worktree()
	if err != nil {
		return err
	}

	found, err := checkForLocalChanges(workTree)
	if err != nil {
		return err
	}

	if found && !force {
		fmt.Println("Local changes were found in the configuration. Use 'octoberctl update -f' to override them.")
		return nil
	}

	if err := pull(repository, octoberConfigDir); err != nil {
		return err
	}

	fmt.Println("October configuration updated.")
	return nil
}

// Checks for local changes in the October Linux configuration.
// Returns found = true if local changes are found and vice versa
func checkForLocalChanges(workTree *git.Worktree) (found bool, err error) {
	status, err := workTree.Status()
	if err != nil {
		return false, err
	}
	for _, value := range status {
		if value.Staging != git.Unmodified || value.Worktree != git.Modified {
			return true, nil
		}
	}
	return false, nil
}

// Pulls the latest version of the October Linux configuration.
// /!\ Overrides local changes on tracked files
func pull(repository *git.Repository, octoberConfigPath string) error {
	if err := repository.Fetch(&git.FetchOptions{Progress: os.Stdout}); err != nil {
		return err
	}

	// Using the cmd version since go-git hard reset removes
	// untracked files.
	cmd := exec.Command("git", "reset", "--hard", "origin/main")
	cmd.Dir = octoberConfigPath

	return cmd.Run()
}
