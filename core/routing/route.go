package igni

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// IMPORTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// STRUCTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

type RouteXML struct {
	XMLName    xml.Name   `xml:"Route"`
	Name       string     `xml:"name,attr"`
	Path       string     `xml:"path,attr"`
	Controller string     `xml:"controller,attr"`
	Handler    string     `xml:"handler,attr"`
	Routes     []RouteXML `xml:"Routes"`
}

type Route struct {
	name       string
	path       string
	controller string
	handler    string
	routes     []*Route
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
	fmt.Printf("Route::handleRequest: %s", request.RequestURI)
	router := GetRouterInstance()

	controller := router.GetHandler(route.handler)
	if controller == nil {
		log.Fatal("Route::handleRequest: controller is nil")
	}

	log.Fatalf("Route::handleRequest: uri='%s'", request.RequestURI)
	controller.HandleRequest(writer, request, route.handler)
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// PUBLIC.METHODS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func (route *Route) Attach(parent string) {
	fmt.Printf("attaching route name='%s', path='%s'", route.name, route.path)
	fmt.Println("")
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
	route.controller = routeXMl.Controller
	for _, subRouteXML := range routeXMl.Routes {
		route.routes = append(route.routes, subRouteXML.fromXML())
	}

	return &route
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
