package main

type AVL struct {
	Value     string
	LeftNode  *AVL
	RightNode *AVL

	ParentNode *AVL

	Height int

	// BalanceFactor is difference between height of LeftNode and RightNode
	// it should be either -1, 0, or 1
	BalanceFactor int
}

// InitializeAvl creates new AVL instance
func InitializeAvl(arr []string) *AVL {
	avl := &AVL{}
	avl.Value = arr[int(len(arr)/2)]

	for _, v := range arr {
		// taking value on the middle
		avl.Insert(v)
	}

	return avl
}

// SetBalanceFactor
func (a *AVL) SetBalanceFactor() {
	if a.ParentNode != nil {
		a.ParentNode.SetBalanceFactor()
	} else {
		// AVL has no parent
		// since, we can calculate a.BalanceFactor
		balanceFactor := a.LeftNode.Height - a.RightNode.Height
		if balanceFactor >= -1 && balanceFactor <= 1 {
			// 			then balanceFactor is ok and we can assign it
			a.BalanceFactor = balanceFactor
		} else {
			// 			if not, then we have a problem
			// 			we need to flip
			//
			//			 root			inserting 1   		root         but balance factor is 2 --->	 changing 		 root
			//			  13								  13														  13
			//			/    \							    /    \													    /     \
			//		   6      20						   6      20												   6       20
			//		  /  \   /  \					   	  /  \   /  \											  	  /  \    /  \
			//		 4   9	15  24						4     9	15  24												3     9	 15  2
			//      /		/						   /	   /												  /	 \      /
			//	   3	   14						  3	  	  14												 1	  4	   14
			//										 /
			//										1
			if balanceFactor > 1 {

			} else {
				// balanceFactor < -1
			}
		}
	}
}

// Search - search operation of AVL using recursion as well
func (a *AVL) Search(value string) *AVL {
	for {
		if a.Value == value {
			return a
		} else if a.Value < value {
			a.LeftNode.Search(value)
		} else {
			// b.Value > value
			a.RightNode.Search(value)
		}
	}
	return nil
}

// InsertToRightRight function inserts value to the most right node in AVL
func (a *AVL) InsertToRightRight(value string) *AVL {
	if a.RightNode != nil {
		a.RightNode.InsertToRightRight(value)
	}

	// if we exited, then we have a.RightNode == nil
	// hence, we can assign it with value from args

	a.RightNode = &AVL{
		Value:         value,
		RightNode:     a,
		BalanceFactor: a.BalanceFactor,
	}

	a.SetBalanceFactor()

	return a.RightNode
}

func (a *AVL) leftLeftRotation() {
	a.ParentNode, a, a.LeftNode = a.RightNode, a.ParentNode, a.LeftNode
}

// subtreeHeight is util function to check nodes height
func (a *AVL) subtreeHeight(node *AVL) int {
	if node == nil {
		return 0
	}

	return node.Height
}

func (a *AVL) calculateAllHeights() {
	// means that node was created but height was not assigned

}

// calculateHeight
func (a *AVL) calculateHeight(avl *AVL) int {
	return avl.LeftNode.Height - avl.RightNode.Height
}

// RightRightRotation performs
//
//		 3 		2
//	    /	   / \
//	   2	  1	  3
//	  /
//	 1
func (a *AVL) RightRightRotation() {
	a, a.LeftNode, a.LeftNode.LeftNode = a.RightNode, a, a.LeftNode
}

func (a *AVL) Insert(value string) *AVL {
	// since we can assign b := &BST{} which is basically nil
	// I've introduced nil check, in order to prevent panic()
	if a == nil {
		return &AVL{Value: value}
	}

	a.Height++

	// if current BST value is less then potentially inserted
	// then, we need to go right
	if a.Value < value {
		a.RightNode = a.RightNode.Insert(value)
	} else if a.Value > value {
		a.LeftNode = a.LeftNode.Insert(value)
	} else {
		return a
	}

	return nil
}
