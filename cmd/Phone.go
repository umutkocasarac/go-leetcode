package main

import (
	"container/list"
	"fmt"
)

func main() {
	number := []int{2, 3, 4}
	fmt.Printf("array is: %v", number)
	letterCombinations(number, len(number))
}

func letterCombinations(number []int, length int) {
	letters := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	queue := list.New()
	list := []string{}
	queue.PushBack("")
	for queue.Len() > 0 {
		item := queue.Back()
		queue.Remove(item)
		value := item.Value.(string)

		if len(value) == length {
			list = append(list, value)
		} else {
			letter := letters[number[len(value)]]
			for i := 0; i < len(letter); i++ {
				queue.PushBack(value + string(letter[i]))
			}
		}
	}
	fmt.Printf("Possible words are %v", list)

}
