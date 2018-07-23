package gonote_test

import (
	"testing"

	"github.com/pokstad/gonote"
)

func TestAllNotes(t *testing.T) {
	allNotes, err := gonote.AllNotes("../", "")
	if err != nil {
		t.Fatalf("unable to parse all notes: %s", err)
	}

	t.Logf("%#v", allNotes)
}
