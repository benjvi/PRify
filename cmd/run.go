package cmd

import (
	"encoding/json"
	"fmt"
	"prify/config"
	"prify/git"
	"prify/diff"
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
	if err != nil {
		log.Fatalf("unable to open %s from current directory: %q", confFile, err)
		os.Exit(1)
	}
	defer file.Close()

	prifyConf, err := config.Parse(file)
	if err != nil {
		log.Fatalf("unable to parse %s: %q", confFile, err)
	}

	// print loaded config
	log.Printf("Loaded %s successfully:", confFile)
	s, _ := json.MarshalIndent(prifyConf, "", "\t")
	fmt.Println(string(s))

	baseBranch, err := git.ResolveBaseBranch(prifyConf.Branch.Base)
	if err != nil {
		log.Fatalf("unable to resolve base branch: %s", err)
	}
	println(baseBranch)

	// make sure that the base branch is up-to-date with upstream
	// for the moment we assume this was done before running prify
	/*upstream := git.ResolveUpstream()
	git.UpdateBaseBranch(baseBranch, upstream)*/

	// find directories with changes in them
	changedDirs, err := diff.FindChangedDirectories()
	if err != nil {
		log.Fatalf("cannot find changed directories: %s", err)
	}
	fmt.Printf("%+v \n", changedDirs)

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