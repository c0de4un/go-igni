package igni

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// IMPORTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// STRUCTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// Application config XML-Struct
type AppConfigXML struct {
	XMLName     xml.Name `xml:"Igni"`
	Debug       bool     `xml:"debug,attr"`
	AppName     string   `xml:"name,attr"`
	Environment string   `xml:"environment,attr"`
	Host        string   `xml:"host,attr"`
	Port        string   `xml:"port,attr"`
}

// Application config struct
type AppConfig struct {
	debug       bool
	name        string
	environment int
	host        string
	port        string
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// GETTERS & SETTERS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func (cfg *AppConfig) IsDebug() bool {
	return cfg.debug
}

func (cfg *AppConfig) GetName() string {
	return cfg.name
}

func (cfg *AppConfig) GetEnvironment() int {
	return cfg.environment
}

func (cfg *AppConfig) GetHost() string {
	return cfg.host
}

func (cfg *AppConfig) GetPort() string {
	return cfg.port
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// METHODS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// Load config
func Load(path string) (*AppConfig, error) {
	var config AppConfig

	if len(path) < 1 {
		return nil, errors.New("AppConfig::load: invalid path")
	}

	xmlFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	xmlBytes, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return nil, err
	}

	var appXml AppConfigXML

	err = xml.Unmarshal(xmlBytes, &appXml)
	if err != nil {
		return nil, err
	}

	config.name = appXml.AppName

	config.debug = appXml.Debug

	if appXml.Environment == "production" {
		config.environment = ENVIRONMENT_PRODUCTION
	} else if appXml.Environment == "development" {
		config.environment = ENVIRONMENT_DEV
	} else if appXml.Environment == "testing" {
		config.environment = ENVIRONMENT_TESTING
	} else {
		return nil, fmt.Errorf("Igni::Load: invalid environment: %s", appXml.Environment)
	}

	config.host = appXml.Host

	return &config, nil
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
