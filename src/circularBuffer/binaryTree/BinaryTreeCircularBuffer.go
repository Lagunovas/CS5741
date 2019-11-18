package binaryTreeCircularBuffer

import (
	binaryTree "github.com/CS5741/src/shared/binaryTree"
)

type BinaryTreeCircularBuffer struct {
	bt       *binaryTree.BinaryTree
	read     int
	write    int
	size     int
	capacity int
}

func NewBinaryTreeCircularBuffer(capacity int) *BinaryTreeCircularBuffer {
	return &BinaryTreeCircularBuffer{bt: binaryTree.NewBinaryTree(), capacity: capacity}
}

func (binaryTreeCircularBuffer *BinaryTreeCircularBuffer) Push(value int) bool {
	if binaryTreeCircularBuffer.CanWrite() {
		if binaryTreeCircularBuffer.write == binaryTreeCircularBuffer.capacity {
			binaryTreeCircularBuffer.write = 0
		}

		binaryTreeCircularBuffer.bt.Remove(binaryTreeCircularBuffer.write + 1)
		binaryTreeCircularBuffer.bt.PushOrder(value, binaryTreeCircularBuffer.write+1)

		binaryTreeCircularBuffer.write++
		binaryTreeCircularBuffer.size++
		return true
	}

	return false
}

func (binaryTreeCircularBuffer *BinaryTreeCircularBuffer) ReadNext() (bool, int) {
	var value int

	if binaryTreeCircularBuffer.HasNext() {
		if binaryTreeCircularBuffer.read == binaryTreeCircularBuffer.capacity {
			binaryTreeCircularBuffer.read = 0
		}

		value = binaryTreeCircularBuffer.bt.NodeAt(binaryTreeCircularBuffer.read + 1).Value()
		binaryTreeCircularBuffer.size--
		binaryTreeCircularBuffer.read++
		return true, value
	}

	return false, value
}

func (binaryTreeCircularBuffer *BinaryTreeCircularBuffer) HasNext() bool {
	return binaryTreeCircularBuffer.size > 0
}

func (binaryTreeCircularBuffer *BinaryTreeCircularBuffer) CanWrite() bool {
	return binaryTreeCircularBuffer.size < binaryTreeCircularBuffer.capacity
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
	binaryTreeCircularBuffer.bt.Clear()
}
