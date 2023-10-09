package main

type AVL struct {
	Value     string
	LeftNode  *AVL
	RightNode *AVL
	Height    int
}
