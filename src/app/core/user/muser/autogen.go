package muser

import (
	"app/core/db"
	"app/core/types"
	"app/core/types/id"
	"app/util/paginate"
	"context"
	"errors"
)

// List fetches a paginated list of entities
func List(ctx context.Context, params *types.PaginationParams) (*types.PaginationResult, error) {
	return paginate.Do(params, func(count uint64, startingAfterID id.IDType) (interface{}, error) {

		out := []*types.User{}
		err := db.SelectContext(
			ctx,
			&out,
			"select * from [User] where ID > :id limit :count",
			map[string]interface{}{
				"id":    startingAfterID,
				"count": count,
			})

		if err != nil {
			return nil, err
		}

		return out, nil
	})
}

// Get fetches a single entity
func Get(ctx context.Context, tenantID id.Tenant, userID id.User) (*types.User, error) {

	out := &types.User{}

	err := db.GetContext(
		ctx,
		out,
		"select * from [User] where TenantID = :tenantID and User = :userID",
		map[string]interface{}{
			"userID":   userID,
			"tenantID": tenantID,
		},
	)

	if err != nil {
		return nil, err
	}

	return out, nil
}

// Create creates a new user
func Create(ctx context.Context, user *types.User) (*types.User, error) {
	return nil, errors.New("todo")
}
