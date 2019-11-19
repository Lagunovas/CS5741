//package concurrentLinkedListCircularBuffer
package main

import(
	"fmt"
	linkedListCircularBuffer "github.com/CS5741/src/circularBuffer/linkedList"
	"sync"
)

type ConcurrentCircularBuffer struct {
	buffer *linkedListCircularBuffer.CircularBuffer
	lock sync.RWMutex
}

func NewConcurrentCircularBuffer() *ConcurrentCircularBuffer {
	buffer := *linkedListCircularBuffer.NewCircularBuffer()
	return &ConcurrentCircularBuffer{buffer: buffer}
}

func (cBuffer *ConcurrentCircularBuffer) Push(value int) bool {
	cBuffer.lock.Lock()
	defer cBuffer.lock.Unlock()
	return cBuffer.buffer.Push(value) 
}

func (cBuffer *ConcurrentCircularBuffer) ReadNext() (bool, int) {
	cBuffer.lock.Lock()
	defer cBuffer.lock.Unlock()
	return cBuffer.buffer.ReadNext()
}

func (cBuffer *ConcurrentCircularBuffer) ToString() string {
	cBuffer.lock.Lock()
	defer cBuffer.lock.Unlock()
	return cBuffer.buffer.ToString()
}

func (cBuffer *ConcurrentCircularBuffer) HasNext() bool {
	cBuffer.lock.Lock()
	defer cBuffer.lock.Unlock()
	return cBuffer.buffer.HasNext() 
}

func (cBuffer *ConcurrentCircularBuffer) Capacity() int {
	cBuffer.lock.Lock()
	defer cBuffer.lock.Unlock()
	return cBuffer.buffer.Capacity()
}