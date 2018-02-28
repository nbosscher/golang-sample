package users

import (
	"app/core/types"
	"app/core/user"
	"app/util"
	"app/util/httprouter"
	"app/util/output"
	"net/http"
)

func InitEndpoints(router httprouter.Router) {
	router.RegisterAuthenticated("GET", "/users", viewUsers)
}

// This method kind of sticks out b/c it's a templated page in an otherwise programmatic API.
// This is just here for proof of concept, probably would only have a handful of these.
func viewUsers(w http.ResponseWriter, r *http.Request) (*output.PageParams, error) {

	pagination := &types.PaginationParams{}

	err := util.DecodeAndValidateQuery(r, pagination)
	if err != nil {
		return nil, err
	}

	list, err := user.List(r.Context(), pagination)
	if err != nil {
		return nil, err
	}

	page := &output.PageParams{
		File: "user-list.html",
		Data: map[string]interface{}{
			"UserList": list,
		},
	}

	return page, nil
}
