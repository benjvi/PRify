package cmd

import (
	"encoding/json"
	"fmt"
	"prify/config"
	"prify/git"
	"github.com/spf13/cobra"
	"log"
	"os"
)

const (
	confFile = "prify.yml"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Run: run,
}


func run(cmd *cobra.Command, args []string) {
	// load configuration from prify.yml
	file, err := os.Open(confFile)
	defer file.Close()
	if err != nil {
		log.Fatalf("unable to open %s from current directory: %q", confFile, err)
		os.Exit(1)
	}

	prifyConf, err := config.Parse(file)
	if err != nil {
		log.Fatalf("unable to parse %s: %q", confFile, err)
	}

	// print loaded config
	log.Printf("Loaded %s successfully:", confFile)
	s, _ := json.MarshalIndent(prifyConf, "", "\t")
	fmt.Println(string(s))

	// find directories with changes in them
	/*baseBranch, changedDirs = findChangedDirectories()

	if len(changedDirs) == 0 {
		print("No changes found")
		return
	}*/

	baseBranch, err := git.ResolveBaseBranch(prifyConf.Branch.Base)
	if err != nil {
		log.Fatalf("unable to resolve base branch: %s", err)
	}
	println(baseBranch)

	// TODO: make sure that the base branch is up-to-date with upstream
	// for the moment we can assume this was done before running prify
	/*upstream := git.ResolveUpstream()
	git.UpdateBaseBranch(baseBranch, upstream)*/

	/*for dir in changedDirs {
		// TODO: resolve templates in config

		// commit to the specified branch
		err = git.BranchAndCommit(dir, prifyConf.Commit, baseBranch, b.Name)
		if err != nil {
		log.Fatalf("error committing changes: %s", err)

		// push branch

		// create PR
	}
	}*/

}