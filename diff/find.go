package diff

import (
	"fmt"
	"os"
	"prify/git"
)


func FindChangedDirectories() ([]string, error) {
	file, err := os.Open(".")
	if err != nil {
		return nil, fmt.Errorf("unable to open PWD", err)
	}
	defer file.Close()

	//TODO: exclude files, should be folders only
	subDirs, err := file.Readdirnames(1000)
	if err != nil {
		return nil, err
	}

	changedDirs := make([]string, len(subDirs))
	for _, dir := range subDirs {
		modified, err := git.IsFolderModified(dir)
		if err != nil {
			return nil, fmt.Errorf("git diff on %q failed: %s", dir, err)
		}
		if modified {
			changedDirs = append(changedDirs, dir)
		}
	}
	return changedDirs, nil
}