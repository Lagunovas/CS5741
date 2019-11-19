package concurrentLinkedListStack

import (
	"sync"

	linkedListStack "github.com/CS5741/src/stack/nonConcurrent/linkedList"
)

type ConcurrentLinkedListStack struct {
	internalStack *linkedListStack.LinkedListStack
	mutex         sync.RWMutex
}

func NewConcurrentLinkedListStack() *ConcurrentLinkedListStack {
	return &ConcurrentLinkedListStack{internalStack: linkedListStack.NewLinkedListStack()}
}

func (concurrentLinkedListStack *ConcurrentLinkedListStack) Push(value int) {
	concurrentLinkedListStack.mutex.Lock()
	defer concurrentLinkedListStack.mutex.Unlock()
	concurrentLinkedListStack.internalStack.Push(value)
}

func (concurrentLinkedListStack *ConcurrentLinkedListStack) Pop() (bool, int) {
	concurrentLinkedListStack.mutex.Lock()
	defer concurrentLinkedListStack.mutex.Unlock()
	return concurrentLinkedListStack.internalStack.Pop()
}

func (concurrentLinkedListStack *ConcurrentLinkedListStack) Peek() (bool, int) {
	concurrentLinkedListStack.mutex.Lock()
	defer concurrentLinkedListStack.mutex.Unlock()
	return concurrentLinkedListStack.internalStack.Peek()
}

func (concurrentLinkedListStack *ConcurrentLinkedListStack) Empty() bool {
	concurrentLinkedListStack.mutex.Lock()
	defer concurrentLinkedListStack.mutex.Unlock()
	return concurrentLinkedListStack.internalStack.Empty()
}

func (concurrentLinkedListStack *ConcurrentLinkedListStack) Size() int {
	concurrentLinkedListStack.mutex.Lock()
	defer concurrentLinkedListStack.mutex.Unlock()
	return concurrentLinkedListStack.internalStack.Size()
}

func (concurrentLinkedListStack *ConcurrentLinkedListStack) Clear() {
	concurrentLinkedListStack.mutex.Lock()
	defer concurrentLinkedListStack.mutex.Unlock()
	concurrentLinkedListStack.internalStack.Clear()
}
