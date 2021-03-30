package template

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"prify/config"
	"github.com/teris-io/shortid"
	"text/template"
)

// when does this error?
var randID, _ = shortid.Generate()

func RenderCommitConf(changesDirRef string, commitConfigTempl config.CommitConfig, baseBranchRef string, targetBranchTempl string) (config.CommitConfig, string, error) {
	extraData := map[string]string{
		"subdir_name": changesDirRef,
		"base_branch": baseBranchRef,
	}

	messageOut, err := render(commitConfigTempl.Message, extraData)
	if err != nil {
		return config.CommitConfig{}, "", fmt.Errorf("couldn't render commit.message: %s", err)
	}

	targetBranchOut, err := render(targetBranchTempl, extraData)
	if err != nil {
		return config.CommitConfig{}, "", fmt.Errorf("couldn't render branch.name: %s", err)
	}
	commitConfigOut := commitConfigTempl
	commitConfigOut.Message = messageOut

	return commitConfigOut, targetBranchOut, nil
}

func render(templateS string, extraData map[string]string) (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("can't get pwd for use in template: %s", err)
	}
	dirName := filepath.Base(pwd)
	data := map[string]string{
		"rand": randID,
		"dir": dirName, // note this should match the prify dir, currently PWD
	}
	for k,v := range extraData {
		data[k] = v
	}
	tmpl, err := template.New("test").Parse(templateS)
	if err != nil {
		return "", fmt.Errorf("error parsing template")
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, data)
	if err != nil {
		return "", fmt.Errorf("error rendering template")
	}
	return buf.String(), nil
}