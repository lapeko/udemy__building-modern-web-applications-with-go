package main

import "fmt"

func RunLoops() {
	fmt.Println("\nStandard loop")
	for i := 0; i < 2; i++ {
		fmt.Println(i)
	}

	fmt.Println("\nLoop from slice")
	names := []string{"Name 1", "Name 2", "Name 3"}

	for index, name := range names {
		fmt.Println(index, name)
	}

	fmt.Println("\nLoop from map")
	languagePopularityMap := make(map[string]int)
	languagePopularityMap["JS"] = 3
	languagePopularityMap["Py"] = 3

	for key, value := range languagePopularityMap {
		fmt.Println(key, value)
	}

	fmt.Println("\nLoop from string")
	line := "Text"

	for _, charByte := range line {
		fmt.Println(charByte)
	}

	fmt.Println("\nLoop from structs")

	type User struct {
		name string
		age  int
	}
	users := []User{{"Andrew", 23}, {"Kate", 21}}

	for _, user := range users {
		fmt.Println("name:", user.name+", age:", user.age)
	}
}
