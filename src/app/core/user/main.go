package user

import (
	"app/core/auth"
	"app/core/auth/permission"
	"app/core/db"
	"app/core/types"
	"app/core/user/muser"
	"context"
	"errors"
)

// List fetches a paginated list of users
func List(ctx context.Context, pagination *types.PaginationParams) (*types.PaginationResult, error) {
	return nil, errors.New("todo")
}

// CreateParams ...
type CreateParams struct {
	User     *types.User `json:"user"`
	Password string      `json:"password"`
}

// Create creates a new user
func Create(ctx context.Context, params *CreateParams) (*types.User, error) {

	creds, err := auth.GetCredentialsWithPermission(ctx, permission.CreateUser)
	if err != nil {
		return nil, err
	}

	var newUser *types.User

	err = db.Tx(ctx, func(ctx context.Context) error {
		newUser, err = muser.Create(ctx, params.User)
		if err != nil {
			return err
		}

		return muser.SetPassword(ctx, params.Password)
	})

	return newUser, err
}
