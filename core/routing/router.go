package igni

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// IMPORTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

import (
	igni_http "github.com/c0de4un/igni/core/http"
)

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// STRUCTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

type Router struct {
	routes      map[string]*Route
	controllers map[int]*igni_http.Controller
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// FIELDS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

var routerInstance *Router

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// GETTERS & SETTERS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// Static instance getters
func GetRouterInstance() *Router {
	return routerInstance
}

func (router *Router) GetHandler(id int) *igni_http.Controller {
	controller, ok := router.controllers[id]
	if !ok {
		return nil
	}

	return controller
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// PUBLIC.METHODS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// New Router instance
func NewRouter() *Router {
	routerInstance = &Router{}
	routerInstance.routes = make(map[string]*Route)
	routerInstance.controllers = make(map[int]*igni_http.Controller)

	return routerInstance
}

func (router *Router) RegisterHandler() {

}

// Load routes
func (router *Router) Load(path string) error {
	config := NewRouterConfig()
	err := config.Load(path)

	return err
}

// Attach handlers
func (router *Router) InitHandlers() {
	for _, route := range router.routes {
		route.Attach("/")
	}
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
