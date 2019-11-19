//package concurrentLinkedListStack
package main

import(
	"fmt"
	linkedListStack "github.com/CS5741/src/stack/linkedList"
	"sync"
)

type ConcurrentStack struct {
	stack *linkedListStack.Stack
	lock sync.RWMutex
}		

func NewConcurrentStack() *ConcurrentStack {
	stack := linkedListStack.NewStack()
	return &ConcurrentStack{stack: stack}
}

func (cstack *ConcurrentStack) Push(value int) { 
	cstack.lock.Lock()
	defer cstack.lock.Unlock()
	return cstack.stack.Push(value)
}

func (cstack *concurrentStack) Pop() (bool, int) {
	cstack.lock.Lock()
	defer cstack.lock.Unlock()
	return cstack.stack.Pop()
}

func (cstack *concurrentStack) Peek() (bool, int) {
	cstack.lock.Lock()
	defer cstack.lock.Unlock()
	return cstack.stack.Peek()
}

func (cstack *concurrentStack) Empty() bool {
	cstack.lock.Lock()
	defer cstack.lock.Unlock()
	return cstack.stack.Empty()
}

func (cstack *concurrentStack) Size() int {
	cstack.lock.Lock()
	defer cstack.lock.Unlock()
	return cstack.stack.Size()
}

func (cstack *concurrentStack) Clear() {
	cstack.lock.Lock()
	defer cstack.lock.Unlock()
	return cstack.stack.Clear()
}