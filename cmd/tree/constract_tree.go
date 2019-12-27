package main

import (
	"fmt"
	"strconv"
)

// constract a tree from inorder and preorder traversal
type node struct {
	value int
	left  *node
	right *node
}

var index int = 0

func main() {
	in := []int{4, 2, 5, 1, 6, 3, 7}
	pre := []int{1, 2, 4, 5, 3, 6, 7}

	root := buildTree(in, pre, 0, len(pre)-1)
	dptFirstPrint(root)
}

func dptFirstPrint(root *node) {
	fmt.Printf("%v\n", root.value)
	if root.left != nil {
		dptFirstPrint(root.left)

	}
	if root.right != nil {
		dptFirstPrint(root.right)
	}

}

func buildTree(in []int, pre []int, start int, end int) *node {
	if start > end {
		return nil
	}
	root := pre[index]
	index = index + 1
	if start == end {
		return &node{value: root}
	}

	index := search(in, root)
	left := buildTree(in, pre, start+1, index-1)
	right := buildTree(in, pre, index+1, end)
	return &node{root, left, right}
}

func search(in []int, value int) int {
	for i := 0; i < len(in); i++ {
		if in[i] == value {
			return i
		}
	}
	panic("couldn't find the item " + strconv.Itoa(value))
}
