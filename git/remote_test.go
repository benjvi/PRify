package git

import (
	"reflect"
	"testing"
)

func TestGetRemoteNames(t *testing.T) {

	got, err := getRemoteNames()
	if err != nil {
		t.Errorf("error getting remotes: %s", err)
	}
	// will fail on a feature branch
	want := []string{"origin", "upstream"}

	if ! reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestResolveRemote(t *testing.T) {
	//repo currently has upstream and origin configured
	// so we should choose origin by default

	got, err := resolveRemote(nil)
	if err != nil {
		t.Errorf("error getting remotes: %s", err)
	}
	// will fail on a feature branch
	want := "origin"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
