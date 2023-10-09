package main

type BST struct {
	Value     string
	LeftNode  *BST
	RightNode *BST
}

// Insert - insertion operation of BST using recursion
func (b *BST) Insert(value string) *BST {
	// since we can assign b := &BST{} which is basically nil
	// I've introduced nil check, in order to prevent panic()
	if b == nil {
		return &BST{Value: value}
	}

	// if current BST value is less then potentially inserted
	// then, we need to go right
	if b.Value < value {
		b.RightNode = b.RightNode.Insert(value)
	} else if b.Value > value {
		b.LeftNode = b.LeftNode.Insert(value)
	} else {
		return b
	}

	return nil
}

// Search - search operation of BST using recursion as well
func (b *BST) Search(value string) *BST {
	for {
		if b.Value == value {
			return b
		} else if b.Value < value {
			b.LeftNode.Search(value)
		} else {
			// b.Value > value
			b.RightNode.Search(value)
		}
	}
	return nil
}
