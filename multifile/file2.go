package main

import (
	"fmt"
)

func main() {
	user := User{
		Name:  "doug",
		Email: "doug.smith@fmr.com",
	}

	fmt.Printf("%v\n", user)

}
