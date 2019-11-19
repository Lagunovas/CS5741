package binaryTreeStack

import (
	binaryTree "github.com/CS5741/src/shared/binaryTree"
)

type BinaryTreeStack struct {
	internalBinaryTree *binaryTree.BinaryTree
}

func NewBinaryTreeStack() *BinaryTreeStack {
	return &BinaryTreeStack{binaryTree.NewBinaryTree()}
}

func (binaryTreeStack *BinaryTreeStack) Push(value int) {
	binaryTreeStack.internalBinaryTree.Push(value)
}

func (binaryTreeStack *BinaryTreeStack) Pop() (bool, int) {
	return binaryTreeStack.internalBinaryTree.Remove(binaryTreeStack.Size())
}

func (binaryTreeStack *BinaryTreeStack) Peek() (bool, int) {
	return binaryTreeStack.internalBinaryTree.Tail()
}

func (binaryTreeStack *BinaryTreeStack) Empty() bool {
	return binaryTreeStack.internalBinaryTree.Empty()
}

func (binaryTreeStack *BinaryTreeStack) Size() int {
	return binaryTreeStack.internalBinaryTree.Size()
}

func (binaryTreeStack *BinaryTreeStack) Clear() {
	binaryTreeStack.internalBinaryTree.Clear()
}
