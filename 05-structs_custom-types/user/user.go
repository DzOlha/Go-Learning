package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// constructor
func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" {
		return nil, errors.New("First name and last name must not be empty")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func NewUser(firstName, lastName, birthDate string) *User {
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}
}

func (u *User) OutputDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) ClearNames() {
	u.firstName = ""
	u.lastName = ""
}
