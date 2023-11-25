package main

import "log"

func RunDataStructures() {
	// var myMap map[string]string = make(map[string]string)
	myMap := make(map[string]string)
	myMap["Cat"] = "Kuzya"

	log.Println(myMap)

	numberSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	log.Println(numberSlice[7])
	log.Println(numberSlice[7:9])
}
