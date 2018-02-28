// Package paginate provides a small helper function.
package paginate

import (
	"app/core/types"
	"app/core/types/id"
	"errors"
)

// Fetcher is a callback that does the actual SQL query.
// Should return an array of entities
type Fetcher = func(count uint64, startingAfterID id.IDType) (interface{}, error)

// Do computes the parameters for fetcher, calls it and builds a paginated result
func Do(params *types.PaginationParams, fetcher Fetcher) (*types.PaginationResult, error) {
	return nil, errors.New("todo")
}
