package concurrentBinaryTreeStack

import (
	"sync"

	binaryTreeStack "github.com/CS5741/src/stack/nonConcurrent/binaryTree"
)

type ConcurrentBinaryTreeStack struct {
	internalStack *binaryTreeStack.BinaryTreeStack
	mutex         sync.RWMutex
}

func NewConcurrentBinaryTreeStack() *ConcurrentBinaryTreeStack {
	return &ConcurrentBinaryTreeStack{internalStack: binaryTreeStack.NewBinaryTreeStack()}
}

func (concurrentBinaryTreeStack *ConcurrentBinaryTreeStack) Push(value int) {
	concurrentBinaryTreeStack.mutex.Lock()
	defer concurrentBinaryTreeStack.mutex.Unlock()
	concurrentBinaryTreeStack.internalStack.Push(value)
}

func (concurrentBinaryTreeStack *ConcurrentBinaryTreeStack) Pop() (bool, int) {
	concurrentBinaryTreeStack.mutex.Lock()
	defer concurrentBinaryTreeStack.mutex.Unlock()
	return concurrentBinaryTreeStack.internalStack.Pop()
}

func (concurrentBinaryTreeStack *ConcurrentBinaryTreeStack) Peek() (bool, int) {
	concurrentBinaryTreeStack.mutex.Lock()
	defer concurrentBinaryTreeStack.mutex.Unlock()
	return concurrentBinaryTreeStack.internalStack.Peek()
}

func (concurrentBinaryTreeStack *ConcurrentBinaryTreeStack) Empty() bool {
	concurrentBinaryTreeStack.mutex.Lock()
	defer concurrentBinaryTreeStack.mutex.Unlock()
	return concurrentBinaryTreeStack.internalStack.Empty()
}

func (concurrentBinaryTreeStack *ConcurrentBinaryTreeStack) Size() int {
	concurrentBinaryTreeStack.mutex.Lock()
	defer concurrentBinaryTreeStack.mutex.Unlock()
	return concurrentBinaryTreeStack.internalStack.Size()
}

func (concurrentBinaryTreeStack *ConcurrentBinaryTreeStack) Clear() {
	concurrentBinaryTreeStack.mutex.Lock()
	defer concurrentBinaryTreeStack.mutex.Unlock()
	concurrentBinaryTreeStack.internalStack.Clear()
}

func (concurrentBinaryTreeStack *ConcurrentBinaryTreeStack) Request() (bool, int) {
	return concurrentBinaryTreeStack.Pop()
}
