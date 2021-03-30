package diff

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestListSubdirsPWD(t *testing.T) {

	// diff folder doesn't currently have subdirs
	got, err := listSubDirs()
	if err != nil {
		t.Errorf("error listing subdirs: %s", err)
	}
	want := []string{}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, wanted %+v", got, want)
	}

}

func TestListSubdirsParent(t *testing.T) {

	err := os.Chdir("../")
	if err != nil {
		t.Errorf("error going to parent folder: %s", err)
	}

	// diff folder doesn't currently have subdirs
	got, err := listSubDirs()
	if err != nil {
		t.Errorf("error listing subdirs: %s", err)
	}
	present := []string{"cmd", "diff"}
	notPresent := []string{"main.go"}

	err = assertContains(got, present)
	if err != nil {
		t.Errorf("contains failed: %s", err)
	}

	err = assertContains(got, notPresent)
	if err == nil {
		t.Errorf("found files %+v in output %+v, should only be directories", notPresent, got )
	}
}

func assertContains(s []string, values []string) error {
	for _, str := range values {
		present := false
		for _, v := range s {
			if v == str {
				present = true
			}
		}
		if !present {
			return fmt.Errorf("Expected value %q not present in %+v", str, s)
		}
	}
	return nil
}