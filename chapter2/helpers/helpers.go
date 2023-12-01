package helpers

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func CalculateSum(num1 int, num2 int) int {
	return num1 + num2
}

func CalculateDivide(a, b float32) (float32, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func PrintText(text *string) {
	fmt.Println(*text)
}

func GetRandomNumberChannel(intChan chan int, size int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	intChan <- r.Intn(size)
}
