package igni

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// IMPORTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

import (
	"log"
	"net/http"
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
	config AppConfig
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

	return app
}

func (app *Igni) Start() error {
	http.HandleFunc("/", app.handleRequest)

	http.ListenAndServe(app.GetUrl(), nil)

	return nil
}

// Router entrypoint
func (app *Igni) handleRequest(response http.ResponseWriter, request *http.Request) {
	message := []byte("Hello Golang Server World !")
	_, err := response.Write(message)
	if err != nil {
		log.Fatalf("Igni::handleRequest: %s", err)
	}
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
