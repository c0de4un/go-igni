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
// STRUCTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

type RouterConfigXML struct {
	XMLName xml.Name   `xml:"Router"`
	Routes  []RouteXML `xml:"Route"`
}

type RouterConfig struct {
	routes map[string]*Route
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// PUBLIC.METHODS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func NewRouterConfig() *RouterConfig {
	var cfg RouterConfig
	cfg.routes = make(map[string]*Route)

	return &cfg
}

func (сfg *RouterConfig) Load(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	var xmlConfig RouterConfigXML
	err = xml.Unmarshal(bytes, &xmlConfig)
	if err != nil {
		return err
	}

	if len(xmlConfig.Routes) < 1 {
		return fmt.Errorf("no routes were found in '%s'", path)
	}

	for _, routeXML := range xmlConfig.Routes {
		if len(routeXML.Name) < 1 {
			return fmt.Errorf("invalid route name")
		}

		сfg.routes[routeXML.Name] = routeXML.fromXML()
	}

	return nil
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
