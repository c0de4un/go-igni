package igni

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// IMPORTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

import (
	"fmt"
	"strconv"

	"github.com/c0de4un/ini"
)

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// CONSTANTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

const APP_CONFIG_PATH string = "data/configs/app.ini"

const ENVIRONMENT_DEV int = 1
const ENVIRONMENT_TESTING int = 2
const ENVIRONMENT_PRODUCTION int = 3

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// STRUCTS.PUBLIC
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// Igni instance
type Igni struct {
	debug       bool
	appName     string
	environment int
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// OVERRIDE: ini.IReaderListener
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func (igni *Igni) OnParam(name string, value string) error {
	if name == "name" {
		igni.appName = value
		return nil
	}

	if name == "environment" {
		var num int
		if value == "production" {
			num = ENVIRONMENT_PRODUCTION
		} else if value == "development" {
			num = ENVIRONMENT_DEV
		} else if value == "testing" {
			num = ENVIRONMENT_TESTING
		} else {
			return fmt.Errorf("invalid environment value: %s", value)
		}

		igni.environment = num
		return nil
	}

	if name == "debug" {
		val, err := strconv.ParseBool(value)
		if err != nil {
			return fmt.Errorf("invalid debug value")
		}

		igni.debug = val
	}

	return nil
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// GETTERS & SETTERS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func (igni *Igni) GetName() string {
	return igni.appName
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// METHODS.PUBLIC
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func NewIgni() Igni {
	var instance Igni
	return instance
}

func (igni *Igni) ReadConfigs() error {
	reader := ini.NewReader()
	return reader.ReadAll(APP_CONFIG_PATH, igni)
}

func (igni *Igni) StartServer() error {
	return nil
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
