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

func createPR(title, body, baseBranch string) (string, error) {
	url,_,err := gh([]string{"pr","create","--title", title, "--body", body, "--base", baseBranch })
	if err != nil {
		return "", fmt.Errorf("error creating pr: %s", err)
	}
	return strings.TrimSpace(url), nil
}

func checkPR(baseBranch string) (bool, error) {
	_,stderr,err := gh([]string{"pr","view", baseBranch})
	if err != nil {
		if strings.Contains(stderr, "no open pull requests found for branch") {
			return false, nil
		} else {
			return false, fmt.Errorf("error checking for existing pr: %s", err)
		}

	}
	//TODO: get URL from the output, which occurs on the line with "View this pull request on GitHub:"
	return true, nil
}

func CreateOrUpdatePR(title, body, baseBranch string) (string, error) {
	prExists, err := checkPR(baseBranch)
	if err != nil {
		return "", err
	}
	if ! prExists {
		return createPR(title, body, baseBranch)
	} else {
		return "pre-existing PR", nil
	}
}