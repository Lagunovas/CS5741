package concurrentArrayStack

import (
	"sync"

	arrayStack "github.com/CS5741/src/stack/nonConcurrent/array"
)

type ConcurrentArrayStack struct {
	internalStack *arrayStack.ArrayStack
	mutex         sync.RWMutex
}

func NewConcurrentArrayStack() *ConcurrentArrayStack {
	return &ConcurrentArrayStack{internalStack: arrayStack.NewArrayStack()}
}

func (concurrentArrayStack *ConcurrentArrayStack) Push(value int) {
	concurrentArrayStack.mutex.Lock()
	defer concurrentArrayStack.mutex.Unlock()
	concurrentArrayStack.internalStack.Push(value)
}

func (concurrentArrayStack *ConcurrentArrayStack) Pop() (bool, int) {
	concurrentArrayStack.mutex.Lock()
	defer concurrentArrayStack.mutex.Unlock()
	return concurrentArrayStack.internalStack.Pop()
}

func (concurrentArrayStack *ConcurrentArrayStack) Peek() (bool, int) {
	concurrentArrayStack.mutex.Lock()
	defer concurrentArrayStack.mutex.Unlock()
	return concurrentArrayStack.internalStack.Peek()
}

func (concurrentArrayStack *ConcurrentArrayStack) Empty() bool {
	concurrentArrayStack.mutex.Lock()
	defer concurrentArrayStack.mutex.Unlock()
	return concurrentArrayStack.internalStack.Empty()
}

func (concurrentArrayStack *ConcurrentArrayStack) Size() int {
	concurrentArrayStack.mutex.Lock()
	defer concurrentArrayStack.mutex.Unlock()
	return concurrentArrayStack.internalStack.Size()
}

func (concurrentArrayStack *ConcurrentArrayStack) Clear() {
	concurrentArrayStack.mutex.Lock()
	defer concurrentArrayStack.mutex.Unlock()
	concurrentArrayStack.internalStack.Clear()
}

func (concurrentArrayStack *ConcurrentArrayStack) Request() (bool, int) {
	return concurrentArrayStack.Pop()
}
