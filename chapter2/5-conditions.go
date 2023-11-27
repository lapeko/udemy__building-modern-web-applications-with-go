package main

import "log"

func RunConditions() {
	num := 5
	isEven := isEven(num)

	if isEven {
		log.Println("var num is even. num: ", num)
	} else {
		log.Println("var num is odd. num: ", num)
	}

	animal := "Cat"

	switch animal {
	case "Cat":
		log.Println("Animal is a cat")
	case "Dog":
		log.Println("Animal is a dog")
	default:
		log.Println("Unknown animal")
	}
}

func isEven(num int) bool {
	return num%2 == 0
}
