package muser

import (
	"app/core/auth"
	"app/core/types/id"
	"context"
	"errors"
)

type SetPasswordParams struct {
	TenantID id.Tenant
	UserID   id.User
	Password string
}

// SetPassword updates the user's password
func SetPassword(ctx context.Context, params *SetPasswordParams) error {
	return errors.New("todo")
}

type GetUserCredentialParams struct {
	Email    string
	Password string
}

// GetUserCredentials looks up the user by username and password and gives the authentication information
func GetUserCredentials(ctx context.Context, params *GetUserCredentialParams) (*auth.Credentials, error) {
	return nil, errors.New("todo")
}
