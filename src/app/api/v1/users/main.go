package users

import (
	"app/core/types"
	"app/core/user"
	"app/util"
	"app/util/httprouter"
	"net/http"
)

func InitEndpoints(router httprouter.Router) {
	router.RegisterAuthenticated("GET", "v1/user/list", listUsers)
	router.RegisterAuthenticated("POST", "v1/user", createUser)
}

func createUser(w http.ResponseWriter, r *http.Request) (*types.User, error) {

	params := &user.CreateParams{}
	err := util.DecodeAndValidateBody(r, params)
	if err != nil {
		return nil, err
	}

	return user.Create(r.Context(), params)
}

func listUsers(w http.ResponseWriter, r *http.Request) (*types.PaginationResult, error) {

	pagination := &types.PaginationParams{}

	err := util.DecodeAndValidateQuery(r, pagination)
	if err != nil {
		return nil, err
	}

	return users.List(r.Context(), pagination)
}
