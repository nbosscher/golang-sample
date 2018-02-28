package v1

import (
	"api/util/httprouter"
	"api/v1/users"
)

func InitEndpoints(router httprouter.Router) {
	users.InitEndpoints(router)
}
