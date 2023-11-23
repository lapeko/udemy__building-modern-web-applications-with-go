package main

import (
	"fmt"
	"time"
)

var someString = "Bla bla"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func RunTypes() {
	fmt.Println(someString)
	someString := "Also bla bla"
	fmt.Println(someString)

	user := User{
		FirstName: "Vitali",
		LastName:  "Lapeka",
	}

	fmt.Println(user)
}
