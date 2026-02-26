package update

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func Update(force bool) error {
	fmt.Println("Updating October configuration...")

	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	repository, err := git.PlainOpen(fmt.Sprintf("%s/october-config", userConfigDir))
	if err != nil {
		return err
	}

	workTree, err := repository.Worktree()
	if err != nil {
		return err
	}

	found, err := checkForChangesGit(workTree)
	if err != nil {
		return err
	}

	if found && !force {
		fmt.Println("The configuration was modified. Use 'octoberctl update -f' to override the changes.")
		return nil
	}

	if err := pull(repository, workTree); err != nil {
		return err
	}

	fmt.Println("October configuration updated.")
	return nil
}

func pull(repository *git.Repository, workTree *git.Worktree) error {
	if err := repository.Fetch(&git.FetchOptions{Progress: os.Stdout}); err != nil {
		return err
	}

	remoteMain, err := repository.Reference(plumbing.NewRemoteReferenceName("origin", "main"), true)
	if err != nil {
		return err
	}

	if err := workTree.Reset(&git.ResetOptions{Mode: git.HardReset, Commit: remoteMain.Hash()}); err != nil {
		return err
	}

	return nil
}

func checkForChangesGit(workTree *git.Worktree) (found bool, err error) {
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
