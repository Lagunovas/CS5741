package binaryTreeCircularBuffer

import (
	binaryTree "github.com/CS5741/src/shared/binaryTree"
)

type BinaryTreeCircularBuffer struct {
	internalBinaryTree *binaryTree.BinaryTree
	read               int
	write              int
	size               int
	capacity           int
}

func NewBinaryTreeCircularBuffer(capacity int) *BinaryTreeCircularBuffer {
	return &BinaryTreeCircularBuffer{internalBinaryTree: binaryTree.NewBinaryTree(), capacity: capacity}
}

func (binaryTreeCircularBuffer *BinaryTreeCircularBuffer) Push(value int) bool {
	if binaryTreeCircularBuffer.canWrite() {
		if binaryTreeCircularBuffer.write == binaryTreeCircularBuffer.capacity {
			binaryTreeCircularBuffer.write = 0
		}

		binaryTreeCircularBuffer.internalBinaryTree.Remove(binaryTreeCircularBuffer.write + 1)
		binaryTreeCircularBuffer.internalBinaryTree.PushOrder(value, binaryTreeCircularBuffer.write+1)

		binaryTreeCircularBuffer.write++
		binaryTreeCircularBuffer.size++
		return true
	}

	return false
}

func (binaryTreeCircularBuffer *BinaryTreeCircularBuffer) HasNext() bool {
	return binaryTreeCircularBuffer.size > 0
}

func (binaryTreeCircularBuffer *BinaryTreeCircularBuffer) ReadNext() (bool, int) {
	if binaryTreeCircularBuffer.HasNext() {
		if binaryTreeCircularBuffer.read == binaryTreeCircularBuffer.capacity {
			binaryTreeCircularBuffer.read = 0
		}

		binaryTreeCircularBuffer.size--
		binaryTreeCircularBuffer.read++

		return binaryTreeCircularBuffer.internalBinaryTree.Get(binaryTreeCircularBuffer.read)
	}

	return false, 0
}

func (binaryTreeCircularBuffer *BinaryTreeCircularBuffer) Capacity() int {
	return binaryTreeCircularBuffer.capacity
}

func (binaryTreeCircularBuffer *BinaryTreeCircularBuffer) Size() int {
	return binaryTreeCircularBuffer.size
}

func (binaryTreeCircularBuffer *BinaryTreeCircularBuffer) Clear() {
	binaryTreeCircularBuffer.read = 0
	binaryTreeCircularBuffer.write = 0
	binaryTreeCircularBuffer.size = 0
	binaryTreeCircularBuffer.internalBinaryTree.Clear()
}

func (binaryTreeCircularBuffer *BinaryTreeCircularBuffer) ToString() string {
	return binaryTreeCircularBuffer.internalBinaryTree.ToString()
}

func (binaryTreeCircularBuffer *BinaryTreeCircularBuffer) canWrite() bool {
	return binaryTreeCircularBuffer.size < binaryTreeCircularBuffer.capacity
}
