package linkedListStack

import (
	linkedList "github.com/CS5741/src/shared/linkedList"
)

type LinkedListStack struct {
	internalLinkedList *linkedList.LinkedList
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{linkedList.NewLinkedList()}
}

func (linkedListStack *LinkedListStack) Push(value int) {
	linkedListStack.internalLinkedList.AddFirst(value)
}

func (linkedListStack *LinkedListStack) Pop() (bool, int) {
	return linkedListStack.internalLinkedList.RemoveFirst()
}

func (linkedListStack *LinkedListStack) Peek() (bool, int) {
	return linkedListStack.internalLinkedList.Head()
}

func (linkedListStack *LinkedListStack) Empty() bool {
	return linkedListStack.internalLinkedList.Empty()
}

func (linkedListStack *LinkedListStack) Size() int {
	return linkedListStack.internalLinkedList.Size()
}

func (linkedListStack *LinkedListStack) Clear() {
	linkedListStack.internalLinkedList.Clear()
}
