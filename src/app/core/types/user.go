package types

import "app/core/types/id"

type User struct {
	ID        id.User `json:"id"`
	Email     string  `json:"email" valid:"email~Invalid email format"`
	FirstName string  `json:"firstName" valid:"required~FirstName is required"`
	LastName  string  `json:"lastName" valid:"required~LastName is required"`
}
