package git

import (
	"fmt"
	"log"
	"prify/config"
)

func MaybePush(pushConf *config.PushConfig, tgtBranch string) error {
	if pushConf != nil {
		// push config was present in the file, so we push
		log.Printf("Push not specified in prify.yml, skipping")
		// lookup remote to push to
		remoteName, err := resolveRemote(pushConf)
		if err != nil {
			return err
		}

		// force is needed because we may have rebased
		_,_,err = git([]string{"push", "-f", remoteName, tgtBranch }, []string{})
		log.Printf("Pushed to remote %s/%s", remoteName, tgtBranch)
		if err != nil {
			return fmt.Errorf("error on git push: %s", err)
		}
	} else {
		log.Printf("Push not specified in prify.yml, skipping")
	}
	return nil
}