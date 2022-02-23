package igni

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// IMPORTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

import "net/http"

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// INTERFACE
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// Request handler contract
type IRequestHandler interface {
	HandleRequest(http.ResponseWriter, *http.Request, string) bool
	GetName() string
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
