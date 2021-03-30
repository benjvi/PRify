package diff

import (
	"fmt"
	"io/ioutil"
	"prify/git"
)


func FindChangedDirectories() ([]string, error) {

	subDirs, err := listSubDirs()
	if err != nil {
		return nil, err
	}

	changedDirs := []string{}
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

func listSubDirs() ([]string, error) {
	// output should be as find . -type d -d 1
	files, err := ioutil.ReadDir("./")
	if err != nil {
		return nil, fmt.Errorf("error listing directories: %s", err)
	}

	subDirs := []string{}
	for _, f := range files {
		// excludes files and current directory
		if (f.IsDir()) {
			subDirs = append(subDirs, f.Name())
		}
	}

	return subDirs, nil
}