package github

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func gh(args []string) (string, string, error) {
	var stdoutBs, stderrBs bytes.Buffer

	cmd := exec.Command("gh", args...)
	cmd.Stdout = &stdoutBs
	cmd.Stderr = &stderrBs

	err := cmd.Run()
	if err != nil {
		return "", "", fmt.Errorf("gh %s: %s (stderr: %s)", args, err, stderrBs.String())
	}

	return stdoutBs.String(), stderrBs.String(), nil
}

func CreatePR(title, body, baseBranch string) (string, error) {
	url,_,err := gh([]string{"pr","create","--title", title, "--body", body, "--base", baseBranch })
	if err != nil {
		return "", fmt.Errorf("error creating pr: %s", err)
	}
	return strings.TrimSpace(url), nil
}