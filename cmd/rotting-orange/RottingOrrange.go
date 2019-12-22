package main

import (
	"container/list"
	"fmt"
)

type node struct {
	row, col int
}

func main() {
	array := [][]uint8{
		{1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1},
	}

	fmt.Printf("array is: %v\n", array)
	fmt.Printf("Minumum days is : %v\n", minumumDays(array))

}

func minumumDays(array [][]uint8) int {
	queue := getRottenNodes(array)
	list := []node{}
	days := 0
	for queue.Len() > 0 {
		item := queue.Front()
		queue.Remove(item)
		node := item.Value.(node)
		list = append(list, rottenNewNodes(node, array)...)

		if queue.Len() == 0 && len(list) > 0 {
			days++
			for i := 0; i < len(list); i++ {
				queue.PushBack(list[i])
			}
			list = list[:0]
		}
	}
	return days
}

func getRottenNodes(array [][]uint8) *list.List {
	queue := list.New()
	row, col := len(array), len(array[0])
	for currentRow := 0; currentRow < row; currentRow++ {
		for currentCol := 0; currentCol < col; currentCol++ {
			if array[currentRow][currentCol] == 1 {
				queue.PushBack(node{currentRow, currentCol})
			}
		}
	}
	return queue
}

func rottenNewNodes(currentNode node, array [][]uint8) []node {
	rows := []int{-1, 0, 0, 1}
	cols := []int{0, -1, 1, 0}
	var nodes = []node{}
	for k := 0; k < len(rows); k++ {
		row, col := currentNode.row+rows[k], currentNode.col+cols[k]
		if isValid(row, col, len(array), len(array[0])) && array[row][col] == 0 {
			array[row][col] = 1
			nodes = append(nodes, node{row, col})
		}
	}
	return nodes
}

func isValid(currentRow int, currentCol int, row int, col int) bool {
	if currentRow >= 0 && currentRow < row && currentCol >= 0 && currentCol < col {
		return true
	}
	return false
}
