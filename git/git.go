package git

import (
	"bytes"
	"fmt"
	"os/exec"
)

// moidification of github.com/vmware-tanzu/carvel-vendir (pkg/vendir/fetch/git/git.go#L75) **APACHE LICENSE
func git(args []string, env []string) (string, string, error) {
	var stdoutBs, stderrBs bytes.Buffer

	cmd := exec.Command("git", args...)
	cmd.Env = env
	cmd.Stdout = &stdoutBs
	cmd.Stderr = &stderrBs

	err := cmd.Run()
	if err != nil {
		return "", "", fmt.Errorf("Git %s: %s (stderr: %s)", args, err, stderrBs.String())
	}

	return stdoutBs.String(), stderrBs.String(), nil
}

