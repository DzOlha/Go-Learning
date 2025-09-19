package admin

import "go-learning/main/05-structs_custom-types/user"

type Admin struct {
	email    string
	password string
	user.User
}

func New(email, password string) *Admin {
	return &Admin{
		email:    email,
		password: password,
		User:     *user.NewUser("ADMIN", "ADMIN", "---"),
	}
}
