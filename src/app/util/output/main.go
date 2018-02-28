package output

import (
	"net/http"
)

// JSON sets the http status code and writes the obj as json to the http response
func JSON(w http.ResponseWriter, statusCode int, obj interface{}) {
	// todo
}

// JSONError works the similar as JSON except it wraps the obj in {error: obj}
func JSONError(w http.ResponseWriter, statusCode int, obj interface{}) {
	// todo
}

// PageParams describes a templated page
type PageParams struct {
	File string
	Data interface{}
}

// Page writes a templated page to the response with the status code given
func Page(w http.ResponseWriter, statusCode int, page *PageParams) {
	// todo
}
