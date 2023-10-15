package main

import (
	"fmt"
	"reflect"
	"time"
)

const (
	ElementLength    = 20
	SearchItemsCount = 5000
	ElementCount     = 2500000
)

func main() {
	var arr []string

	for i := 0; i < ElementCount; i++ {
		arr = append(arr, generateRandomString(ElementLength))
	}

	// avl := InitializeSingleAVL(generateRandomString(ElementLength))
	bst := InitializeBST(arr)
	avl := Insert(nil, generateRandomString(ElementLength))

	avlInsertionStartTime := time.Now().UnixMilli()
	for i := 0; i < ElementCount; i++ {
		avl = Insert(avl, generateRandomString(ElementLength))
	}

	notFoundElementsAVl := []string{}
	notFoundElementsBST := []string{}

	avlDurationTime := time.Now().UnixMilli() - avlInsertionStartTime

	bstInsertionStartTime := time.Now().UnixMilli()
	for i := 0; i < ElementCount; i++ {
		bst.Insert(generateRandomString(ElementLength))
	}

	bstDurationTime := time.Now().UnixMilli() - bstInsertionStartTime

	fmt.Printf(`
			Insertion time of AVL and BST.
			AVL: %d
			BST: %d
	-------------------------------------------
		`,
		avlDurationTime,
		bstDurationTime)

	var stringsToFind []string

	for i := 0; i < SearchItemsCount; i++ {
		stringsToFind = append(stringsToFind, generateRandomString(ElementLength))
	}

	for _, v := range stringsToFind {
		Insert(avl, v)
		bst.Insert(v)
	}

	avlSearchStart := time.Now().UnixMilli()

	for _, v := range stringsToFind {
		ok := Search(avl, v)
		if !ok {
			notFoundElementsAVl = append(notFoundElementsAVl, v)
			fmt.Println("not found bst")
		}
	}

	avlSearchDuration := time.Now().UnixMilli() - avlSearchStart

	bstSearchStart := time.Now().UnixMilli()

	for _, v := range stringsToFind {
		ok := Search(avl, v)
		if !ok {
			notFoundElementsBST = append(notFoundElementsBST, v)
			fmt.Println("not found avl")
		}
	}

	bstSearchDuration := time.Now().UnixMilli() - bstSearchStart

	fmt.Println(reflect.DeepEqual(notFoundElementsBST, notFoundElementsAVl))

	fmt.Printf(`
				Searching time in
				AVL: %d
				BST: %d

			`,
		avlSearchDuration,
		bstSearchDuration)

}
