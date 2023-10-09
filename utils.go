package main

import "math/rand"

// generateRandomString is utility function to fill AVL and BST with random string values
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// sortStrArray basically returns sorted array of strings
func sortStrArray(arr []string) []string {
	for _, v := range arr {
		for i := 0; i <= len(arr)-1; i++ {
			if v < arr[i] {
				v, arr[i] = arr[i], v
			}
		}
	}

	return arr
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
