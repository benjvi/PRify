package cmd

import (
	"prify/config"
	"prify/git"
	"prify/diff"
	"github.com/spf13/cobra"
	"log"
	"os"
	"prify/template"
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

	userSuppliedConf, err := config.Parse(file)
	if err != nil {
		log.Fatalf("unable to parse %s: %q", confFile, err)
	}


	log.Printf("Loaded %s successfully", confFile)
	// print loaded config
	/*log.Println("Contents:")
	s, _ := json.MarshalIndent(userSuppliedConf, "", "\t")
	fmt.Println(string(s))*/

	baseBranch, err := git.ResolveBaseBranch(userSuppliedConf.Branch.Base)
	if err != nil {
		log.Fatalf("unable to resolve base branch: %s", err)
	}

	// IMPROVEMENT: make sure that the base branch is up-to-date with upstream
	// for the moment we assume this was done before running prify
	/*upstream := git.ResolveUpstream()
	git.UpdateBaseBranch(baseBranch, upstream)*/

	// find directories with changes in them
	changedDirs, err := diff.FindChangedDirectories()
	if err != nil {
		log.Fatalf("cannot find changed directories: %s", err)
	}

	for _, dir := range changedDirs {
		// late binding of the templates - we render right before the action so we have all the info
		log.Printf("Rendering template for %q", dir)
		commitR, tgtBranchR, err := template.RenderCommitConf(dir, userSuppliedConf.Commit, baseBranch, userSuppliedConf.Branch.Name)
		if err != nil {
			log.Fatalf("commit conf rendering for %q failed: %s", dir, err)
		}

		log.Printf("Committing changes for %q", dir)
		err = git.BranchAndCommit(dir, commitR, baseBranch, tgtBranchR)
		if err != nil {
			log.Fatalf("error committing changes: %s", err)
		}

		// push branch
		err = git.MaybePush(userSuppliedConf.Push, tgtBranchR)
		if err != nil {
			log.Fatalf("error pushing branch: %s", err)
		}

		// create PR
	}

}