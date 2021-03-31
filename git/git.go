package git

import (
	"bytes"
	"fmt"
	"os/exec"
)

// moidification of github.com/vmware-tanzu/carvel-vendir (pkg/vendir/fetch/git/git.go#L75) **APACHE LICENSE
func git(args []string) (string, string, error) {
	var stdoutBs, stderrBs bytes.Buffer

	cmd := exec.Command("git", args...)
	cmd.Stdout = &stdoutBs
	cmd.Stderr = &stderrBs

	err := cmd.Run()
	if err != nil {
		return "", "", fmt.Errorf("git %s: %s (stderr: %s)", args, err, stderrBs.String())
	}

	return stdoutBs.String(), stderrBs.String(), nil
}

func IsFolderModified(folder string) (bool, error) {
	// TODO: unit test
	// ls-files is the way to reliably find changed files
	// see https://stackoverflow.com/a/35375409 and docs https://git-scm.com/docs/git-ls-files
	changed_files, _, err := git([]string{"ls-files", "--others", "--modified", "--deleted", "--exclude-standard", folder})
	if err != nil {
		return false, err
	}

	if len(changed_files) > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
