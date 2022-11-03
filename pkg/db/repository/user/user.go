package user

import (
	"Api/pkg/db/connect"
	"Api/pkg/models/users"
)

var Repository users.UserRepository

func init() {
	Repository = users.ProvideUserRepository(connect.DB)
}
