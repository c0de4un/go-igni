package igni

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// IMPORTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

import (
	"testing"
)

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// UNITS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

/// Configs loading test
func TestReadConfigs(t *testing.T) {
	app := NewIgni()
	err := app.ReadConfigs()
	if err != nil {
		t.Errorf("Igni::TestReadConfigs: %s", err)
	}
}

/// Server Start test

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
