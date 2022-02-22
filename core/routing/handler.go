package igni

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// IMPORTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

import "net/http"

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// INTERFACE
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// Request Handler
type IRequestHandler interface {
	HandleRequest(writer http.ResponseWriter, request *http.Request, target string)
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
