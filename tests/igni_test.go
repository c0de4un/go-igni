package igni

import (
	"testing"

	"github.com/c0de4un/igni"
)

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// IMPORTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// UNITS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

/// Configs loading test
func TestLoading(t *testing.T) {
	app := igni.NewIgni("..")
	err := app.Load()
	if err != nil {
		t.Errorf("Igni::Load: failed with error: %s", err)
	}
}

/// Server Start test

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
