package main

// package binarytree

import "fmt"

type BinaryTree struct {
	count int
	order int
	root  *Node
}

type Node struct {
	order int
	value int
	left  *Node
	right *Node
}

func (node *Node) IsRoot() bool {
	return node.order == 0
}

func (node *Node) IsLeaf() bool {
	return node.left == node.right && node.right == nil
}

func NewNode() *Node {
	return &Node{0, 0, nil, nil}
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{0, 0, nil}
}

func (binaryTree *BinaryTree) Empty() bool {
	return binaryTree.root == nil
}

func (binaryTree *BinaryTree) Size() int {
	return binaryTree.count
}

func (binaryTree *BinaryTree) Push(value int) bool {
	var node *Node = binaryTree.root

	if binaryTree.Empty() {
		binaryTree.root = NewNode()
		node = binaryTree.root
	} else {
		for {
			if value < node.value {
				if node.left == nil {
					node.left = NewNode()
					node = node.left
					break
				} else {
					node = node.left
				}
			} else if value > node.value {
				if node.right == nil {
					node.right = NewNode()
					node = node.right
					break
				} else {
					node = node.right
				}
			} else {
				return false
			}
		}
	}

	node.value = value
	node.order = binaryTree.order

	binaryTree.count++
	binaryTree.order++
	return true
}

// Level Order Traversal (BFS)
func (BinaryTree *BinaryTree) Print() string {
	return "asd"
}

func main() {
	var binaryTree *BinaryTree = NewBinaryTree()
	fmt.Print(binaryTree.Push(5))
	fmt.Print(binaryTree.Push(6))
	fmt.Print(binaryTree.Push(4))
	fmt.Print(binaryTree.count)
}
