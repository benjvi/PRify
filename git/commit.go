package git

import (
	"fmt"
	"prify/config"
)

func BranchAndCommit(changesDir string, c config.CommitConfig, baseBranch string, targetBranch string) error {

	// TODO: calculate if branch exists and make sure its up-to-date
	branchExists := false
	if branchExists {
		git([]string{"checkout", targetBranch}, []string{})
	} else {
		// make sure we're on the base branch before creating new branch
		git([]string{"checkout", baseBranch}, []string{})
		git([]string{"checkout","-b", targetBranch}, []string{})
	}

	git([]string{"add", changesDir}, []string{})

	git([]string{"commit", "-m", c.Message, fmt.Sprintf("%s <%s>", c.Author, c.Email) }, []string{})

	if baseBranch != targetBranch {
		git([]string{"rebase", "-X", "theirs"}, []string{})
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
	return branch, nil
}

func fetchBranch(branch string) (bool, error) {
	exists := true
	return exists, nil
}

func checkoutBranch(exists bool) {
}