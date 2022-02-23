package igni

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// IMPORTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

import (
	"fmt"
	"log"
	"net/http"
)

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// STRUCTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

type Controller struct {
	name string
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// METHODS.PUBLIC
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func (controller *Controller) GetName() string {
	return controller.name
}

func (controller *Controller) HandleRequest(response http.ResponseWriter, request *http.Request, handler string) bool {
	fmt.Printf("Controller::HandleRequest: handler=%s; uri=%s;", handler, request.RequestURI)
	if handler == "index" {
		controller.index(response, request)
		return true
	}

	return false
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// METHODS.PRIVATE
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func (controller *Controller) index(response http.ResponseWriter, request *http.Request) {
	writeResponse("Hello Golang Server World !", response)
}

func writeResponse(message string, response http.ResponseWriter) {
	bytes := []byte(message)
	_, err := response.Write(bytes)
	if err != nil {
		log.Fatalf("Controller::writeResponse: %s", err)
	}
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
