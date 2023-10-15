package main

import (
	"fmt"
	"math/rand"
	sort "sort"
)

// generateRandomString is utility function to fill AVL and BST with random string values
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// sortStrArray basically returns sorted array of strings (bubble sort)

func sortStrArray(arr []string) {
	sort.Strings(arr)
}

// min utils function returns minimal value from array of string
func min(array []string) string {
	min := array[0]
	for _, v := range array {
		if v < min {
			min = v
		}
	}

	return min
}

// max is utility function to check which of two values is greater
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func printAVLTree(node *AVL, indent int) {
	if node == nil {
		return
	}

	// print the current node
	fmt.Printf("%s%s\n", getIndentString(indent), node.Value)

	// print left subtree
	fmt.Printf("%sLeft:\n", getIndentString(indent+1))
	printAVLTree(node.LeftNode, indent+2)

	// print right subtree
	fmt.Printf("%sRight:\n", getIndentString(indent+1))
	printAVLTree(node.RightNode, indent+2)
}

func getIndentString(indent int) string {
	result := ""
	for i := 0; i < indent; i++ {
		result += "  " // Use 2 spaces for each level of indentation
	}
	return result
}
