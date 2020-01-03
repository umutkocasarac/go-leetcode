package main

import "fmt"

func main() {
	coins := []int{2, 3, 5, 10}
	fmt.Printf("alternative ways: %v", findTotalAlternatives(coins, 12))
}

func findTotalAlternatives(coins []int, value int) int {
	alternatives := make([]int, value+1)
	alternatives[0] = 1
	for _, coin := range coins {
		for i := coin; i <= value; i++ {
			remain := i - coin
			alternatives[i] = alternatives[i] + alternatives[remain]
		}
	}
	return alternatives[value]
}
