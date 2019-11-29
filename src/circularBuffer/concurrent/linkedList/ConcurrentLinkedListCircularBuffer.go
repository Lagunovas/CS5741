package concurrentLinkedListCircularBuffer

import (
	"sync"

	linkedListCircularBuffer "github.com/CS5741/src/circularBuffer/nonConcurrent/linkedList"
)

type ConcurrentLinkedListCircularBuffer struct {
	internalBuffer *linkedListCircularBuffer.LinkedListCircularBuffer
	mutex          sync.RWMutex
}

func NewConcurrentCircularBuffer(capacity int) *ConcurrentLinkedListCircularBuffer {
	return &ConcurrentLinkedListCircularBuffer{internalBuffer: linkedListCircularBuffer.NewLinkedListCircularBuffer(capacity)}
}

func (concurrentLinkedListCircularBuffer *ConcurrentLinkedListCircularBuffer) Push(value int) bool {
	concurrentLinkedListCircularBuffer.mutex.Lock()
	defer concurrentLinkedListCircularBuffer.mutex.Unlock()
	return concurrentLinkedListCircularBuffer.internalBuffer.Push(value)
}

func (concurrentLinkedListCircularBuffer *ConcurrentLinkedListCircularBuffer) HasNext() bool {
	concurrentLinkedListCircularBuffer.mutex.Lock()
	defer concurrentLinkedListCircularBuffer.mutex.Unlock()
	return concurrentLinkedListCircularBuffer.internalBuffer.HasNext()
}

func (concurrentLinkedListCircularBuffer *ConcurrentLinkedListCircularBuffer) ReadNext() (bool, int) {
	concurrentLinkedListCircularBuffer.mutex.Lock()
	defer concurrentLinkedListCircularBuffer.mutex.Unlock()
	return concurrentLinkedListCircularBuffer.internalBuffer.ReadNext()
}

func (concurrentLinkedListCircularBuffer *ConcurrentLinkedListCircularBuffer) Capacity() int {
	concurrentLinkedListCircularBuffer.mutex.Lock()
	defer concurrentLinkedListCircularBuffer.mutex.Unlock()
	return concurrentLinkedListCircularBuffer.internalBuffer.Capacity()
}

func (concurrentLinkedListCircularBuffer *ConcurrentLinkedListCircularBuffer) ToString() string {
	concurrentLinkedListCircularBuffer.mutex.Lock()
	defer concurrentLinkedListCircularBuffer.mutex.Unlock()
	return concurrentLinkedListCircularBuffer.internalBuffer.ToString()
}

func (concurrentLinkedListCircularBuffer *ConcurrentLinkedListCircularBuffer) Size() int {
	return concurrentLinkedListCircularBuffer.internalBuffer.Size()
}

func (concurrentLinkedListCircularBuffer *ConcurrentLinkedListCircularBuffer) Clear() {
	concurrentLinkedListCircularBuffer.internalBuffer.Clear()
}

func (concurrentLinkedListCircularBuffer *ConcurrentLinkedListCircularBuffer) Request() (bool, int) {
	return concurrentLinkedListCircularBuffer.ReadNext()
}
