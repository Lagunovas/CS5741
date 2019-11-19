package concurrentArrayCircularBuffer

import (
	"sync"

	arrayCircularBuffer "github.com/CS5741/src/circularBuffer/nonConcurrent/array"
)

type ConcurrentArrayCircularBuffer struct {
	internalBuffer *arrayCircularBuffer.ArrayCircularBuffer
	mutex          sync.RWMutex
}

func NewConcurrentArrayCircularBuffer(capacity int) *ConcurrentArrayCircularBuffer {
	return &ConcurrentArrayCircularBuffer{internalBuffer: arrayCircularBuffer.NewArayCircularBuffer(capacity)}
}

func (concurrentArrayCircularBuffer *ConcurrentArrayCircularBuffer) Push(value int) bool {
	concurrentArrayCircularBuffer.mutex.Lock()
	defer concurrentArrayCircularBuffer.mutex.Unlock()
	return concurrentArrayCircularBuffer.internalBuffer.Push(value)
}

func (concurrentArrayCircularBuffer *ConcurrentArrayCircularBuffer) HasNext() bool {
	concurrentArrayCircularBuffer.mutex.Lock()
	defer concurrentArrayCircularBuffer.mutex.Unlock()
	return concurrentArrayCircularBuffer.internalBuffer.HasNext()
}

func (concurrentArrayCircularBuffer *ConcurrentArrayCircularBuffer) ReadNext() (bool, int) {
	concurrentArrayCircularBuffer.mutex.Lock()
	defer concurrentArrayCircularBuffer.mutex.Unlock()
	return concurrentArrayCircularBuffer.internalBuffer.ReadNext()
}

func (concurrentArrayCircularBuffer *ConcurrentArrayCircularBuffer) Capacity() int {
	concurrentArrayCircularBuffer.mutex.Lock()
	defer concurrentArrayCircularBuffer.mutex.Unlock()
	return concurrentArrayCircularBuffer.internalBuffer.Capacity()
}

func (concurrentArrayCircularBuffer *ConcurrentArrayCircularBuffer) ToString() string {
	return "NOT IMPLEMENTED"
}
