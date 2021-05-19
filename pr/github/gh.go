package github

import (
	"bytes"
	"fmt"
	"log"
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

func checkPRExists(baseBranch string) (bool, error) {
	stdout,stderr,err := gh([]string{"pr","view", baseBranch, "--json",
		"state,url", "-q", "select(.state == \"OPEN\")" })
	if err != nil {
			return false, fmt.Errorf("error checking for existing pr: %s", err)
	}
	log.Printf("Command `gh pr view %s` stdout message: %q", baseBranch, stdout)
	log.Printf("Command `gh pr view %s` stderr message: %q", baseBranch, stderr)
	//TODO: get URL from the output, which occurs on the line with "View this pull request on GitHub:"
	if strings.Contains(stdout,"url") {
		return true, nil
	}
	return false, nil
}

func CreateOrUpdatePR(title, body, baseBranch, targetBranch string) (string, error) {
	prExists, err := checkPRExists(targetBranch)
	if err != nil {
		return "", err
	}
	if ! prExists {
		return createPR(title, body, baseBranch)
	} else {
		return "pre-existing PR", nil
	}
}