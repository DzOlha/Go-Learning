package structs

import (
	"fmt"
	"go-learning/main/05-structs_custom-types/admin"
	"go-learning/main/05-structs_custom-types/user"
)

func RunTheApp() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	u, err := user.New(firstName, lastName, birthdate)

	administrator := admin.New("email", "password")
	administrator.OutputDetails()

	if err != nil {
		fmt.Println(err)
		return
	}

	// ... do something awesome with that gathered data!

	u.OutputDetails()
	u.ClearNames()
	u.OutputDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
