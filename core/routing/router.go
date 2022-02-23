package igni

import (
	"fmt"
	"log"
)

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// IMPORTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// STRUCTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

type Router struct {
	routes   map[string]*Route
	handlers map[string]IRequestHandler
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// FIELDS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

var routerInstance *Router

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// METHODS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func GetRouterInstance() *Router {
	return routerInstance
}

func (router *Router) GetHandler(id string) IRequestHandler {
	handler, ok := router.handlers[id]
	if !ok {
		log.Fatalf("Router::GetHandler: no handler were found for %s", id)
	}

	return handler
}

func NewRouter() *Router {
	routerInstance = &Router{}
	routerInstance.routes = make(map[string]*Route)
	routerInstance.handlers = make(map[string]IRequestHandler)

	return routerInstance
}

func (router *Router) RegisterHandler(handler IRequestHandler) {
	if len(handler.GetName()) < 1 {
		log.Fatal("Router::RegisterHandler: invalid handler name")
	}

	fmt.Printf("Router::RegisterHandler: %s", handler.GetName())
	fmt.Println("")

	router.handlers[handler.GetName()] = handler
}

func (router *Router) Load(path string) error {
	routes, err := LoadRouterConfig(path)
	if err != nil {
		return err
	}

	router.routes = routes
	return nil
}

func (router *Router) InitHandlers() {
	if len(router.handlers) < 1 {
		log.Fatal("Router::InitHandlers: no handlers were added")
	}
	if len(router.routes) < 1 {
		log.Fatal("Router::InitHandlers: no routes were loaded")
	}

	for _, route := range router.routes {
		route.Attach("/")
	}
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
