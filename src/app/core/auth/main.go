package auth

import (
	"app/core/auth/permission"
	"app/core/types/id"
	"app/util/sandboxerr"
	"context"
	"errors"
	"net/http"
)

// Handler checks for proper authorization of http request and
// returns un-authorized error if it fails
func Handler(handler http.HandlerFunc) http.Handler {
	return nil
}

// Credentials ...
type Credentials struct {
	TenantID    id.Tenant
	LocationID  id.Location
	UserID      id.User
	permissions map[permission.Perm]bool
}

// HasPermission checks whether the credential has permission to something
func (c *Credentials) HasPermission(perm permission.Perm) bool {
	// todo
	return false
}

// ErrNoPermission ...
var ErrNoPermission = errors.New("no-permission")

// GetCredentialsWithPermission gets the credentials from ctx and checks if
// the user has permission.
//
// may return ErrNoPermission, ErrNoCredentials
func GetCredentialsWithPermission(ctx context.Context, perm permission.Perm) (*Credentials, error) {
	return nil, sandboxerr.Wrap(errors.New("todo"))
}

// ErrNoCredentials ...
var ErrNoCredentials = errors.New("no credentials")

// GetCredentials gets the credentials embedded in the ctx
// returns ErrNoCredentials if not found
func GetCredentials(ctx context.Context) (*Credentials, error) {
	return nil, sandboxerr.Wrap(errors.New("todo"))
}
