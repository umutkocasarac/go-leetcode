package main

import "fmt"

const ALPHABET_SIZE = 26

// http://www.code2succeed.com/golang-insert-and-search-trie/
type TrieNode struct {
	children   [ALPHABET_SIZE]*TrieNode
	endOfWords bool
}

func newNode() *TrieNode {
	node := &TrieNode{}
	node.endOfWords = false

	for i := 0; i < ALPHABET_SIZE; i++ {
		node.children[i] = nil
	}

	return node
}

func main() {
	words := []string{"a", "and", "an", "go", "golang", "man", "mango"}
	root := newNode()

	for i := 0; i < len(words); i++ {
		insert(root, words[i])
	}

	fmt.Println("contains words [a]: ", search(root, "a"))
	fmt.Println("contains words [mango]: ", search(root, "mango"))
	fmt.Println("contains words [lang]: ", search(root, "lang"))
}

func insert(root *TrieNode, key string) {
	temp := root
	for i := 0; i < len(key); i++ {
		index := key[i] - 'a'
		if temp.children[index] == nil {
			temp.children[index] = newNode()
		}
		temp = temp.children[index]
	}
	temp.endOfWords = true
}

func search(root *TrieNode, key string) bool {
	temp := root
	for i := 0; i < len(key); i++ {
		index := key[i] - 'a'
		if temp.children[index] == nil {
			return false
		}
		temp = temp.children[index]
	}
	return temp != nil && temp.endOfWords
}
