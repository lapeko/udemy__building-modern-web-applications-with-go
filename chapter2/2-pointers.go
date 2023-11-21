package main

import "fmt"

func RunPointers() {
	name := "Marta"

	modifyName(&name)

	fmt.Print("Modified name is: \"")
	fmt.Print(name)
	fmt.Println("\"")
}

func modifyName(name *string) {
	*name = ""
}
