package igni

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// IMPORTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// CONSTANTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

const APP_CONFIG_PATH string = "data/configs/app.xml"

const ENVIRONMENT_DEV int = 1
const ENVIRONMENT_TESTING int = 2
const ENVIRONMENT_PRODUCTION int = 3

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// STRUCTS.PUBLIC
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

type IgniXML struct {
	XMLName     xml.Name `xml:"Igni"`
	Debug       bool     `xml:"debug,attr"`
	AppName     string   `xml:"name,attr"`
	Environment string   `xml:"environment,attr"`
}

// Igni instance
type Igni struct {
	debug        bool
	appName      string
	environment  int
	pathModifier string
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// GETTERS & SETTERS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func (app *Igni) IsDebug() bool {
	return app.debug
}

func (app *Igni) GetAppName() string {
	return app.appName
}

func (app *Igni) GetEnvironment() int {
	return app.environment
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// PUBLIC.METHODS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

/// Creates new Igni instance
/// path - path to app root dir
/// leave empty, if called from app dir already
func NewIgni(path string) Igni {
	var app Igni
	app.pathModifier = path

	return app
}

/// Load params from config-file
func (app *Igni) Load() error {
	xmlFile, err := os.Open(app.pathModifier + "/" + APP_CONFIG_PATH)
	if err != nil {
		return err
	}
	defer xmlFile.Close()

	xmlBytes, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return err
	}

	var appXml IgniXML

	err = xml.Unmarshal(xmlBytes, &appXml)
	if err != nil {
		return nil
	}

	app.appName = appXml.AppName

	app.debug = appXml.Debug

	if appXml.Environment == "production" {
		app.environment = ENVIRONMENT_PRODUCTION
	} else if appXml.Environment == "development" {
		app.environment = ENVIRONMENT_DEV
	} else if appXml.Environment == "testing" {
		app.environment = ENVIRONMENT_TESTING
	} else {
		return fmt.Errorf("Igni::Load: invalid environment: %s", appXml.Environment)
	}

	return nil
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
