package git

import (
	"testing"
)

func TestResolveBaseBranch(t *testing.T) {

	got, err := ResolveBaseBranch(nil)
	if err != nil {
		t.Errorf("error resolving branch: %s", err)
	}
	// will fail on a feature branch
	want := "main"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
