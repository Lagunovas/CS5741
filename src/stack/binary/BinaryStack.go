package binaryStack

import (
	binaryTree "github.com/CS5741/src/shared"
)

type BinaryStack struct {
	bt *binaryTree.BinaryTree
}

func NewBinaryStack() *BinaryStack {
	return &BinaryStack{binaryTree.NewBinaryTree()}
}

func (binaryStack *BinaryStack) Push(value int) {
	binaryStack.bt.Push(value)
}

func (binaryStack *BinaryStack) Pop() (bool, int) {
	return binaryStack.bt.Remove(binaryStack.Size())
}

func (binaryStack *BinaryStack) Peek() (bool, int) {
	topElement := binaryStack.bt.Tail()

	if topElement != nil {
		return true, topElement.Value()
	}

	return false, 0
}

func (binaryStack *BinaryStack) Empty() bool {
	return binaryStack.bt.Empty()
}

func (binaryStack *BinaryStack) Size() int {
	return binaryStack.bt.Size()
}

func (binaryStack *BinaryStack) Clear() {
	binaryStack.bt.Clear()
}
