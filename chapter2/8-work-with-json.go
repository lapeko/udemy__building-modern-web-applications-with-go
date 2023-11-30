package main

import (
	"encoding/json"
	"fmt"
)

func RunJsonExample() {
	jsonExample := `[{"name": "Name 1", "age": 20},{"name": "Name 2", "age": 30}]`

	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	var users []User

	err := json.Unmarshal([]byte(jsonExample), &users)

	if err != nil {
		fmt.Println("JSON parse error", err)
	}

	fmt.Printf("users: %v\n", users)

	newJson, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		fmt.Println("Marshalling error occurred", err)
	}
	fmt.Println("New JSON from slice:\n" + string(newJson))
}
