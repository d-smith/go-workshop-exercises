package main

import (
	"fmt"
)

func main() {
	user := User{
		Name:  "doug",
		Email: "doug.smith.mail@gmail.com",
	}

	fmt.Printf("%v\n", user)

}
