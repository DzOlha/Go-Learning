package structs

import "fmt"

type customString string

func (str *customString) log() {
	fmt.Println(*str)
}

func AliasExample() {
	var name customString = "Olha"
	name.log()
}
