package git

import (
	"fmt"
	"prify/config"
	"strings"
)

func resolveRemote(userSuppliedConf *config.PushConfig) (string, error) {
	// TODO: if remote is explicitly specified in pushconfig, use that

	// else if git repo has a default remote, use that
	remoteNames, err := getRemoteNames()
	if err != nil {
		return "", err
	}

	if len(remoteNames) == 1 {
		return remoteNames[0], nil
	} else if len(remoteNames) > 1 && contains(remoteNames, "origin") {
		return "origin", nil
	}

	return "", fmt.Errorf("couldn't determine a remote to push to from remotes %+v", remoteNames)
}

func getRemoteNames() ([]string, error) {
	stdOut,_,err := git([]string{"remote"})
	if err != nil {
		return nil, fmt.Errorf("error running git remote: %s", err)
	}
	return strings.Split(strings.TrimSpace(stdOut),"\n"), nil
}

func contains(s []string, value string) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}
	return false
}

/*
type remote struct {
	name string
	url string
	remoteType string
}*/

/*func fetchRemoteBranch() {
	// fetch branch or fail
	remoteBranchExists := false
	if remote != nil {
		remoteBranchExists, err := fetchBranch(b.Name, remote)
		if err != nil {
			return err
		}


	}
}*/