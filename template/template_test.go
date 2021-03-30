package template

import (
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

func TestPlainRender(t *testing.T) {

	template := "simple: {{.one}} {{.two}}"
	values := map[string]string{
		"one": "foo",
		"two": "bar",
	}
	got, err := render(template, values)
	if err != nil {
		t.Errorf("render threw error: %s", err)
	}
	want := "simple: foo bar"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestRenderWithRand(t *testing.T)  {
	template := "{{.rand}}"
	values := map[string]string{
		"one": "foo",
		"two": "bar",
	}
	got, err := render(template, values)
	if err != nil {
		t.Errorf("render threw error: %s", err)
	}
	//shortid should be 9 char long
	if len(got) != 9 {
		t.Errorf("wrong length of %q, got %d, wanted %d", got, len(got), 9)
	}
	//shortid should only contain URL-friendly chars, see https://github.com/teris-io/shortid/blob/771a37caa5cf0c81f585d7b6df4dfc77e0615b5c/shortid.go#L85
	match := regexp.MustCompile(`^[A-Za-z0-9_\-]+$`).MatchString
	if !match(got) {
		t.Errorf("%q has unexpected chars\n", got)
	}
}

func TestRenderWithPWD(t *testing.T)  {
	template := "{{.dir}}"
	values := map[string]string{
		"one": "foo",
		"two": "bar",
	}
	got, err := render(template, values)
	if err != nil {
		t.Errorf("render threw error: %s", err)
	}

	// should be able to stat returned dir value
	// its a relative path so we need to get the parent path
	pwd, err := os.Getwd()
	if err != nil {
		t.Errorf("couldn't get PWD to verify dir: %s", err)
	}
	dir := filepath.Dir(pwd)
	if _, err := os.Stat(filepath.Join(dir, got)); err != nil {
		t.Errorf("opening returned path gave an error: %s", err)
	}
}
