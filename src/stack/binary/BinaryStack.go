package main

import (
	"fmt"

	binaryTree "github.com/CS5741/src/shared"
)

type BinaryStack struct {
	bt    *binaryTree.BinaryTree
	order int
}

func NewBinaryStack() *BinaryStack {
	return &BinaryStack{binaryTree.NewBinaryTree(), 0}
}

func (binaryStack *BinaryStack) Push(value int) {
	// arrayStack.items = append(arrayStack.items, value)
	// arrayStack.size++
	binaryStack.bt.Push(value)
}

func (binaryStack *BinaryStack) Pop() (bool, int) {
	var removed, value = binaryStack.bt.Remove(binaryStack.order)

	if removed {
		binaryStack.order++
	}

	return removed, value
}

// func (binaryStack *BinaryStack) Peek() int {
// 	var item int

// 	if !arrayStack.Empty() {
// 		item = arrayStack.items[arrayStack.size-1]
// 	}

// 	return item
// }

// func (binaryStack *BinaryStack) Empty() bool {
// 	return arrayStack.size == 0
// }

func (binaryStack *BinaryStack) Size() int {
	return binaryStack.bt.Size()
}

// func (binaryStack *BinaryStack) Clear() {
// 	arrayStack.items = nil
// 	arrayStack.size = 0
// }

func main() {
	var bs *BinaryStack = NewBinaryStack()
	bs.Push(3)
	bs.Push(2)
	bs.Push(1)

	fmt.Printf("1 size: %d\n", bs.Size())

	var r, v = bs.Pop()

	fmt.Printf("r: %v, v: %v\n", r, v)
	fmt.Printf("2 size: %d\n", bs.Size())
	r, v = bs.Pop()

	fmt.Printf("r: %v, v: %v\n", r, v)
	fmt.Printf("2 size: %d\n", bs.Size())
	r, v = bs.Pop()

	fmt.Printf("r: %v, v: %v\n", r, v)
	fmt.Printf("2 size: %d\n", bs.Size())
}
