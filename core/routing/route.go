package igni

import "encoding/xml"

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
	routes  []*Route `xml:"Routes"`
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

func (route *Route) GetHandler() int {
	return route.handler
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// PUBLIC.METHODS
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
