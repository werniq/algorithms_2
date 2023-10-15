package main

type BST struct {
	Value     string
	LeftNode  *BST
	RightNode *BST
}

func InitializeBST(arr []string) *BST {
	if len(arr) == 0 {
		return nil
	}

	bst := &BST{
		Value: arr[0],
	}

	for i := 1; i < len(arr); i++ {
		bst.Insert(arr[i])
	}

	return bst
}

func (b *BST) Insert(value string) *BST {
	if value < b.Value {
		if b.LeftNode == nil {
			b.LeftNode = &BST{Value: value}
		} else {
			b.LeftNode.Insert(value)
		}
	} else {
		if b.RightNode == nil {
			b.RightNode = &BST{Value: value}
		} else {
			b.RightNode.Insert(value)
		}
	}
	return b
}

func (b *BST) Search(value string) bool {
	if b == nil {
		return false
	}

	if value == b.Value {
		return true
	} else if value < b.Value {
		return b.LeftNode.Search(value)
	} else {
		return b.RightNode.Search(value)
	}
}
