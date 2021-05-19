package github

import (
	"fmt"
	"testing"
)

func CreateTestPRManager() PRManager {
	c := testCmd{}
	pm := defaultPRManager{cmd: &c}
	return &pm
}

type testCmd struct {}

func (c *testCmd) run(args []string) (string, string, error) {
	// checkPRExists func: run([]string{"pr","view", baseBranch, "--json",
	//		"state,url", "-q", "select(.state == \"OPEN\")" })
	if args[0] == "pr" && args[1] == "view" {
		if args[2] == "new-branch" {
			return "","",nil
		} else if args[2] == "already-existing-open-pr-branch" {
			return "{\"state\":\"OPEN\",\"url\":\"https://github.com/some-org/some-repo/pull/42\"}", "", nil
		}

	}

	// createPR func: run([]string{"pr","create","--title", title, "--body", body, "--base", baseBranch })
	// returns error on fail, we don't care about stdout/stderr
	// will fail if title or baseBranch already matches a PR in the repo (TODO: double-check)
	if args[0] == "pr" && args[1] == "create" {
		if args[7] == "main" {
			return "made-up PR creation succeeded message", "", nil
		}
	}

	return "", "", fmt.Errorf("received unexpected arguments %v", args)
}

func TestBranchAlreadyExists(t *testing.T) {
	m := CreateTestPRManager()
	_, err := m.CreateOrUpdatePR("new user-specified title", "user-specified description", "main", "already-existing-open-pr-branch")
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestNewBranch(t *testing.T) {
	m := CreateTestPRManager()
	_, err := m.CreateOrUpdatePR("new user-specified title", "user-specified description", "main", "new-branch")
	if err != nil {
		t.Fatalf(err.Error())
	}
}

/* TODO: it seems github doesn't care if multiple titles have the same title, double-check
func TestTitleAlreadyExists(t *testing.T) {
	m := CreateTestPRManager()
	_, err := m.CreateOrUpdatePR("already-existing user-specified title", "user-specified description", "main", "new-branch")
	if err != nil {
		t.Fatalf(err.Error())
	}}

func TestTitleAndBranchAlreadyExists(t *testing.T) {
	m := CreateTestPRManager()
	_, err := m.CreateOrUpdatePR("already-existing user-specified title", "user-specified description", "main", "new-branch")
	if err != nil {
		t.Fatalf(err.Error())
	}}
*/