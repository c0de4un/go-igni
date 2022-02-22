package igni

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// IMPORTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

import "net/http"

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// STRUCT
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

type Controller struct {
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// METHODS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func NewController() *Controller {
	return &Controller{}
}

// Handle request
// target - name of last segment
func (controller *Controller) HandleRequest(writer http.ResponseWriter, request *http.Request) {
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
