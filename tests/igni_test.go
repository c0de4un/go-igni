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
func TestAppLoading(t *testing.T) {
	cfg, err := igni.Load("../" + igni.APP_CONFIG_PATH)
	if err != nil {
		t.Errorf("failed to load app config, error: %s", err)
		return
	}

	if cfg.GetName() != "Igni" {
		t.Errorf("Igni::Load: invalid name: %s", cfg.GetName())
		return
	}
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
