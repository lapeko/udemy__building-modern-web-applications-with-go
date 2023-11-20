package main

import "fmt"

func main() {
	fmt.Println(123)

	var whatToSay string
	var num int

	whatToSay = "Goodbye, cruel world !"
	num = 7

	fmt.Println(whatToSay)
	fmt.Println("I is set to", num)

	whatWasSaid, otherThing := saySomething()
	fmt.Println(whatWasSaid, otherThing)
}

func saySomething() (string, string) {
	return "Something", "other thing"
}
