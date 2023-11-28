package main

import "fmt"

type Dog struct {
	weight int
}

type Gorilla struct {
	color string
}

func RunInterfaces() {
	dog := Dog{12}
	gorilla := Gorilla{"green"}

	printAnimal(dog)
	printAnimal(gorilla)
}

func (dog Dog) Says() string {
	return "Woof"
}
func (dog Dog) Breed() string {
	return "Dog"
}

func (dog Gorilla) Says() string {
	return "Uuhhaay"
}
func (dog Gorilla) Breed() string {
	return "Gorilla"
}

type Animal interface {
	Says() string
	Breed() string
}

func printAnimal(animal Animal) {
	fmt.Println("The animal is:", animal.Breed(), "it says: \""+animal.Says()+"\"")
}
