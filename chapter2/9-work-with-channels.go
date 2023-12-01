package main

import (
	"chapter2/helpers"
	"fmt"
	"strconv"
)

func RunWorkWithChannels() {
	sum := strconv.Itoa(helpers.CalculateSum(2, 3))
	helpers.PrintText(&sum)

	randomChannel := make(chan int)
	defer close(randomChannel)

	go helpers.GetRandomNumberChannel(randomChannel, 100)

	num := <-randomChannel
	fmt.Println(num)
}
