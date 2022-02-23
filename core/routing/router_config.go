package igni

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// IMPORTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// STRUCTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

type RouterConfigXML struct {
	XMLName xml.Name   `xml:"Router"`
	Routes  []RouteXML `xml:"Route"`
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// PUBLIC.METHODS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func LoadRouterConfig(path string) (map[string]*Route, error) {
	routes := make(map[string]*Route)

	file, err := os.Open(path)
	if err != nil {
		return routes, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return routes, err
	}

	var xmlConfig RouterConfigXML
	err = xml.Unmarshal(bytes, &xmlConfig)
	if err != nil {
		return routes, err
	}

	if len(xmlConfig.Routes) < 1 {
		return routes, fmt.Errorf("no routes were found in '%s'", path)
	}

	for _, routeXML := range xmlConfig.Routes {
		if len(routeXML.Name) < 1 {
			return routes, fmt.Errorf("invalid route name")
		}

		routes[routeXML.Name] = routeXML.fromXML()
	}

	return routes, nil
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
