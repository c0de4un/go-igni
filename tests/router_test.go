package igni

import (
	"testing"

	igni_core "github.com/c0de4un/igni"
	igni_router "github.com/c0de4un/igni/core/routing"
)

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// IMPORTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// UNITS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// Routes loading test
func TestRouterLoading(t *testing.T) {
	configs := igni_router.NewRouterConfig()

	err := configs.Load("../" + igni_core.APP_ROUTER_CONFIG_PATH)
	if err != nil {
		t.Errorf("failed: RouterConfig::Load: %s", err)
	}
}

// @TODO: Routes search test

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
