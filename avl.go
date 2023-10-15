package main

type AVL struct {
	Value     string
	LeftNode  *AVL
	RightNode *AVL

	Height int
}

func InitializeAvl(arr []string) *AVL {

	avl := &AVL{}
	avl.Value = arr[len(arr)/2]

	for i, v := range arr {
		if i != len(arr)/2 {
			avl = Insert(avl, v)
		}
	}

	return avl
}

func InitializeSingleAVL(value string) *AVL {
	return &AVL{
		LeftNode:  &AVL{},
		RightNode: &AVL{},
		Value:     value,
		Height:    1,
	}
}

func Search(avl *AVL, value string) bool {
	if avl == nil {
		return false
	}

	if value == avl.Value {
		return true
	} else if value > avl.Value {
		return Search(avl.RightNode, value)
	} else {
		return Search(avl.LeftNode, value)
	}
}

func calcBalanceFactor(a *AVL) int {
	if a == nil {
		return 0
	}

	return getNodeHeight(a.LeftNode) - getNodeHeight(a.RightNode)
}

func getNodeHeight(avl *AVL) int {
	if avl == nil {
		return 0
	}

	return avl.Height
}

func Insert(a *AVL, value string) *AVL {
	if a == nil {
		return InitializeSingleAVL(value)
	}

	if value > a.Value {
		a.RightNode = Insert(a.RightNode, value)
	} else if value < a.Value {
		a.LeftNode = Insert(a.LeftNode, value)
	} else {
		return a
	}

	a.Height = 1 + max(getNodeHeight(a.LeftNode), getNodeHeight(a.RightNode))
	bF := calcBalanceFactor(a)

	if bF > 1 {
		if value < a.LeftNode.Value {
			return rightRotate(a)
		} else if value > a.RightNode.Value {
			a.LeftNode = leftRotate(a.LeftNode)
			return rightRotate(a)
		}
	}

	if bF < -1 {
		if value > a.RightNode.Value {
			return leftRotate(a)
		} else if value < a.RightNode.Value {
			a.RightNode = rightRotate(a.RightNode)
			return leftRotate(a)
		}
	}

	return a
}

func rightRotate(a *AVL) *AVL {
	newRoot := a.LeftNode
	newRightNode := newRoot.RightNode
	newRoot.RightNode = a

	a.LeftNode = newRightNode
	a.Height = max(getNodeHeight(a.LeftNode), getNodeHeight(a.RightNode)) + 1

	newRoot.Height = max(getNodeHeight(newRoot.LeftNode), getNodeHeight(newRoot.RightNode)) + 1

	return newRoot
}

func leftRotate(a *AVL) *AVL {
	newRoot := a.RightNode
	newLeftNode := newRoot.LeftNode
	newRoot.LeftNode = a

	a.RightNode = newLeftNode
	a.Height = max(getNodeHeight(a.LeftNode), getNodeHeight(a.RightNode)) + 1

	newRoot.Height = max(getNodeHeight(newRoot.LeftNode), getNodeHeight(newRoot.RightNode)) + 1

	return newRoot
}
