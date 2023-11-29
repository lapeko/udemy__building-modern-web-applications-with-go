package helpers

import (
	"fmt"
	"math/rand"
	"time"
)

func CalculateSum(num1 int, num2 int) int {
	return num1 + num2
}

func PrintText(text *string) {
	fmt.Println(*text)
}

func GetRandomNumberChannel(intChan chan int, size int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	intChan <- r.Intn(size)
}
