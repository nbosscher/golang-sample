package util

import (
	"errors"
	"net/http"
)

// DecodeAndValidateQuery decodes query parameters into the dest object
// and validates dest using https://github.com/asaskevich/govalidator
//
// dest should be a pointer
func DecodeAndValidateQuery(r *http.Request, dest interface{}) error {
	return errors.New("todo")
}

// DecodeAndValidateBody decodes the request body into the dest object
// and validates dest using https://github.com/asaskevich/govalidator
//
// Body format is determined by the content-type header. This
// method supports JSON, URL-encoded and MultiPart Forms.
//
// TODO: determin how to handle multipart form files
//
// dest should be a pointer.
func DecodeAndValidateBody(r *http.Request, dest interface{}) error {
	return errors.New("todo")
}
