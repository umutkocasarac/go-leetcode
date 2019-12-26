package main

import (
	"fmt"
)

func main() {
	value := 12
	fmt.Printf("fibonacci: %v\n", fibonacci(value))
	fmt.Printf("fibonacci: %v", fibonacciDp(value))
}

func fibonacci(value int) int {
	if value < 2 {
		return value
	} else {
		return fibonacci(value-1) + fibonacci(value-2)
	}
}

func fibonacciDp(value int) int {
	arr := make([]int, value+1)
	arr[0] = 0
	arr[1] = 1
	for i := 2; i <= value; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[value]
}
