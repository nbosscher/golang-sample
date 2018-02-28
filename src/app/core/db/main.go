// Package db is an interface to sqlx that puts the query in a transaction scope
// if the context has a transaction.
package db

import (
	"app/util/sandboxerr"
	"context"
	"errors"
)

// Tx makes calls within the callback func to db (e.g. db.SelectContext())
// execute within a transaction scope
func Tx(ctx context.Context, callback func(ctx context.Context) error) error {
	// add something to ctx
	// call callback with new ctx
	return sandboxerr.Wrap(errors.New("todo"))
}

// SelectContext is a mock of https://github.com/jmoiron/sqlx
func SelectContext(ctx context.Context, list interface{}, query string, params interface{}) error {
	return sandboxerr.Wrap(errors.New("todo"))
}

// GetContext is a mock of https://github.com/jmoiron/sqlx
func GetContext(ctx context.Context, output interface{}, query string, params interface{}) error {
	return sandboxerr.Wrap(errors.New("todo"))
}
