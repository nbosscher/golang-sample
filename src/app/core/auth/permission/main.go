package permission

type Perm string

const (
	CreateUser Perm = "user.create"
	UpdateUser Perm = "user.update"
	DeleteUser Perm = "user.delete"
	ListUsers  Perm = "users.list"
)
