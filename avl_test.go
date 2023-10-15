package main

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestAVL_Insert(t *testing.T) {
	cases := []string{
		"ddd", "bbb", "ccc", "aaa", "asdasd",
	}

	avl := InitializeAvl(cases)

	avl = Insert(avl, "xxxx")

	avl = Insert(avl, "mmmmmm")

	avl = Insert(avl, "lllll")

	l := getNodeHeight(avl.LeftNode)
	r := getNodeHeight(avl.RightNode)

	ok := l-r == 0 || l-r == 1 || l-r == -1

	printAVLTree(avl, 2)

	assert.Equal(t, ok, true)
}
