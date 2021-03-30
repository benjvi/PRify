package git

import (
	"fmt"
	"log"
	"prify/config"
	"strings"
)

func BranchAndCommit(changesDir string, c config.CommitConfig, baseBranch string, targetBranch string) error {

	// TODO: calculate if branch exists and make sure its up-to-date
	err := checkoutBranch(baseBranch, targetBranch)
	if err != nil  {
		return err
	}

	_,_,err = git([]string{"add", changesDir}, []string{})
	if err != nil {
		return fmt.Errorf("error on git add %s: %s", changesDir, err)
	}

	//TODO fmt.Sprintf("%s <%s>", *c.Author, *c.Email) }, for setting author
	// ignore any pre-commit hooks with `-n`, because tools may not be available
	_,_,err = git([]string{"commit", "-n", "-m", c.Message}, []string{})
	if err != nil {
		return fmt.Errorf("error on git commit: %s", err)
	}

	// rebase branch in case its old and out-of date
	// TODO: review rebase strategy
	if baseBranch != targetBranch {
		_,_,err := git([]string{"rebase", "-X", "theirs", "--autostash", baseBranch}, []string{})
		if err != nil {
			return fmt.Errorf("error on git rebase: %s", err)
		}
	}

	return nil
}

func ResolveBaseBranch(userSpecifiedBranch *string) (string, error) {
	//TODO: validate userSpecifiedBranch?
	if userSpecifiedBranch != nil {
		return *userSpecifiedBranch, nil
	}

	branch, _, err := git([]string{"rev-parse","--abbrev-ref","HEAD"}, []string{})
	if err != nil {
		return "", fmt.Errorf("error reading current branch: %s", err)
	}
	return strings.TrimSpace(branch), nil
}

func checkoutBranch(baseBranch, targetBranch string) error {
	branchExists := false
	if branchExists {
		log.Printf("Using existing branch %q", targetBranch)
		_,_,err := git([]string{"checkout", targetBranch}, []string{})
		if err != nil {
			return fmt.Errorf("error on git checkout %s: %s", targetBranch, err)
		}
	} else {
		// make sure we're on the base branch before creating new branch
		log.Printf("Creating new branch %q", targetBranch)
		_,_,err := git([]string{"checkout", baseBranch}, []string{})
		if err != nil {
			return fmt.Errorf("error on checkout %s: %s", baseBranch, err)
		}
		_,_,err = git([]string{"checkout","-b", targetBranch}, []string{})
		if err != nil {
			return fmt.Errorf("error on git checkout  -b %s: %s", targetBranch, err)
		}
	}
	return nil
}