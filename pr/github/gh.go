package github

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type Cmd interface {
	run(args []string) (string, string, error)
}
type defaultCmd struct {}

type PRManager interface {
	createPR(title, body, baseBranch string) (string, error)
	checkPRExists(baseBranch string) (bool, error)
	CreateOrUpdatePR(title, body, baseBranch, targetBranch string) (string, error)
}

type defaultPRManager struct {
	cmd Cmd
}

func NewPRManager() PRManager {
	c := defaultCmd{}
	pm := defaultPRManager{cmd: &c}
	return &pm
}

func (c *defaultCmd) run(args []string) (string, string, error) {
	var stdoutBs, stderrBs bytes.Buffer

	cmd := exec.Command("run", args...)
	cmd.Stdout = &stdoutBs
	cmd.Stderr = &stderrBs

	err := cmd.Run()
	if err != nil {
		return "", "", fmt.Errorf("run %s: %s (stderr: %s)", args, err, stderrBs.String())
	}

	return stdoutBs.String(), stderrBs.String(), nil
}

func (m *defaultPRManager) createPR(title, body, baseBranch string) (string, error) {
	// we have the target branch checked-out, so just need to specify the base branch for merge
	url,_,err := m.cmd.run([]string{"pr","create","--title", title, "--body", body, "--base", baseBranch })
	if err != nil {
		return "", fmt.Errorf("error creating pr: %s", err)
	}
	return strings.TrimSpace(url), nil
}

func (m *defaultPRManager) checkPRExists(baseBranch string) (bool, error) {
	stdout,stderr,err := m.cmd.run([]string{"pr","view", baseBranch, "--json",
		"state,url", "-q", "select(.state == \"OPEN\")" })
	if err != nil {
			return false, fmt.Errorf("error checking for existing pr: %s", err)
	}
	log.Printf("Command `run pr view %s` stdout message: %q", baseBranch, stdout)
	log.Printf("Command `run pr view %s` stderr message: %q", baseBranch, stderr)
	//TODO: get URL from the output
	if strings.Contains(stdout,"url") {
		return true, nil
	}
	return false, nil
}

func (m *defaultPRManager) CreateOrUpdatePR(title, body, baseBranch, targetBranch string) (string, error) {
	prExists, err := m.checkPRExists(targetBranch)
	if err != nil {
		return "", err
	}
	if ! prExists {
		return m.createPR(title, body, baseBranch)
	} else {
		return "pre-existing PR", nil
	}
}