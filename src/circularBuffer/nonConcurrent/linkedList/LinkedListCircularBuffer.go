package linkedListCircularBuffer

import (
	linkedList "github.com/CS5741/src/shared/linkedList"
)

type LinkedListCircularBuffer struct {
	internalLinkedList *linkedList.LinkedList
	read               int
	write              int
	size               int
	capacity           int
}

func NewLinkedListCircularBuffer(capacity int) *LinkedListCircularBuffer {
	return &LinkedListCircularBuffer{linkedList.NewLinkedList(), 0, 0, 0, capacity}
}

func (linkedListCircularBuffer *LinkedListCircularBuffer) Push(value int) bool {
	if linkedListCircularBuffer.size < linkedListCircularBuffer.capacity {
		if linkedListCircularBuffer.write == linkedListCircularBuffer.capacity {
			linkedListCircularBuffer.write = 0
		}

		if !linkedListCircularBuffer.internalLinkedList.Set(linkedListCircularBuffer.write, value) {
			linkedListCircularBuffer.internalLinkedList.AddLast(value)
		}

		linkedListCircularBuffer.size++
		linkedListCircularBuffer.write++
		return true
	}

	return false
}

func (linkedListCircularBuffer *LinkedListCircularBuffer) HasNext() bool {
	return linkedListCircularBuffer.size > 0
}

func (linkedListCircularBuffer *LinkedListCircularBuffer) ReadNext() (bool, int) {
	if linkedListCircularBuffer.HasNext() {
		if linkedListCircularBuffer.read == linkedListCircularBuffer.capacity {
			linkedListCircularBuffer.read = 0
		}

		linkedListCircularBuffer.size--
		linkedListCircularBuffer.read++

		return linkedListCircularBuffer.internalLinkedList.Get(linkedListCircularBuffer.read - 1)
	}

	return false, 0
}

func (linkedListCircularBuffer *LinkedListCircularBuffer) Capacity() int {
	return linkedListCircularBuffer.capacity
}

func (linkedListCircularBuffer *LinkedListCircularBuffer) Size() int {
	return linkedListCircularBuffer.size
}

func (linkedListCircularBuffer *LinkedListCircularBuffer) Clear() {
	linkedListCircularBuffer.internalLinkedList.Clear()
}

func (linkedListCircularBuffer *LinkedListCircularBuffer) ToString() string {
	return linkedListCircularBuffer.internalLinkedList.ToString()
}
