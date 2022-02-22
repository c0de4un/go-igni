package igni

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// IMPORTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

type RouteXML struct {
	XMLName xml.Name   `xml:"Route"`
	Name    string     `xml:"name,attr"`
	Path    string     `xml:"path,attr"`
	Handler int        `xml:"handler,attr"`
	Routes  []RouteXML `xml:"Routes"`
}

type Route struct {
	name    string
	path    string
	handler int
	routes  []*Route
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -\
// GETTERS & SETTERS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func (route *Route) GetName() string {
	return route.name
}

func (route *Route) GetPath() string {
	return route.path
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// OVERRIDE: http.Handler
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func (route *Route) handleRequest(writer http.ResponseWriter, request *http.Request) {
	router := GetRouterInstance()

	controller := router.GetHandler(route.handler)

	controller.HandleRequest(writer, request)
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// PUBLIC.METHODS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// Attach route and its children to muxer
func (route *Route) Attach(parent string) {
	fmt.Printf("attaching route %s", route.name)
	path := parent + route.path
	http.HandleFunc(path, route.handleRequest)

	for _, child := range route.routes {
		child.Attach(path + "/")
	}
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// PRIVATE.METHODS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func (routeXMl *RouteXML) fromXML() *Route {
	var route Route

	route.name = routeXMl.Name
	route.path = routeXMl.Path
	route.handler = routeXMl.Handler
	for _, subRouteXML := range routeXMl.Routes {
		route.routes = append(route.routes, subRouteXML.fromXML())
	}

	return &route
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
