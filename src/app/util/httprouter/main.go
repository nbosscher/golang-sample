package httprouter

import (
	"net/http"
)

// Router builds up a route list on startup and
// multiplexes requests to matched routes.
type Router interface {
	// RegisterAuthenticated adds an authenticated endpoint.
	//
	// handler should be of type
	// func(w http.ResponseWriter, r *http.Request) (interface{}, error)
	// handler output is formatted through
	RegisterAuthenticated(method string, path string, handler interface{})

	http.Handler
}
