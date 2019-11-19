package concurrentBinaryTreeCircularBuffer

import (
	"sync"

	binaryTreeCircularBuffer "github.com/CS5741/src/circularBuffer/nonConcurrent/binaryTree"
)

type ConcurrentBinaryTreeCircularBuffer struct {
	internalBuffer *binaryTreeCircularBuffer.BinaryTreeCircularBuffer
	mutex          sync.RWMutex
}

func NewConcurrentBinaryTreeCircularBuffer(capacity int) *ConcurrentBinaryTreeCircularBuffer {
	return &ConcurrentBinaryTreeCircularBuffer{internalBuffer: binaryTreeCircularBuffer.NewBinaryTreeCircularBuffer(capacity)}
}

func (concurrentBinaryTreeCircularBuffer *ConcurrentBinaryTreeCircularBuffer) Push(value int) bool {
	concurrentBinaryTreeCircularBuffer.mutex.Lock()
	defer concurrentBinaryTreeCircularBuffer.mutex.Unlock()
	return concurrentBinaryTreeCircularBuffer.internalBuffer.Push(value)
}

func (concurrentBinaryTreeCircularBuffer *ConcurrentBinaryTreeCircularBuffer) HasNext() bool {
	concurrentBinaryTreeCircularBuffer.mutex.Lock()
	defer concurrentBinaryTreeCircularBuffer.mutex.Unlock()
	return concurrentBinaryTreeCircularBuffer.internalBuffer.HasNext()
}

func (concurrentBinaryTreeCircularBuffer *ConcurrentBinaryTreeCircularBuffer) ReadNext() (bool, int) {
	concurrentBinaryTreeCircularBuffer.mutex.Lock()
	defer concurrentBinaryTreeCircularBuffer.mutex.Unlock()
	return concurrentBinaryTreeCircularBuffer.internalBuffer.ReadNext()
}

func (concurrentBinaryTreeCircularBuffer *ConcurrentBinaryTreeCircularBuffer) Capacity() int {
	concurrentBinaryTreeCircularBuffer.mutex.Lock()
	defer concurrentBinaryTreeCircularBuffer.mutex.Unlock()
	return concurrentBinaryTreeCircularBuffer.internalBuffer.Capacity()
}

func (concurrentBinaryTreeCircularBuffer *ConcurrentBinaryTreeCircularBuffer) ToString() string {
	return "NOT IMPLEMENTED"
}
