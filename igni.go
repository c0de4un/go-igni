package igni

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// IMPORTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

import (
	"fmt"
	"net/http"

	igni_router "github.com/c0de4un/igni/core/routing"
)

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// CONSTANTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

const APP_CONFIG_PATH string = "data/configs/app.xml"
const APP_ROUTER_CONFIG_PATH string = "data/configs/router.xml"

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// STRUCTS.PUBLIC
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// Igni instance
type Igni struct {
	pathModifier string
	config       AppConfig
	router       igni_router.Router
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// GETTERS & SETTERS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func (app *Igni) GetUrl() string {
	return app.config.host + ":" + app.config.port
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// PUBLIC.METHODS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

/// Creates new Igni instance
/// path - path to app root dir
/// leave empty, if called from app dir already
func New(path string) Igni {
	var app Igni
	app.pathModifier = path

	return app
}

func (app *Igni) Start() error {
	fmt.Println("main: loading routing")
	err := app.router.Load(app.pathModifier + APP_ROUTER_CONFIG_PATH)
	if err != nil {
		return err
	}
	app.router.InitHandlers()

	fmt.Println("main: starting http server")
	http.ListenAndServe(app.GetUrl(), nil)

	return nil
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
