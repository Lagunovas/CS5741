package linkedList

import (
	"bytes"
	"strconv"
)

type LinkedListNode struct {
	value int
	next  *LinkedListNode
}

func NewLinkedListNode(value int) *LinkedListNode {
	return &LinkedListNode{value: value}
}

type LinkedList struct {
	head *LinkedListNode
	tail *LinkedListNode
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{nil, nil, 0}
}

func (linkedList *LinkedList) AddFirst(value int) {
	node := NewLinkedListNode(value)

	if linkedList.Empty() { //if the linked list has no elements in it
		linkedList.head = node
		linkedList.tail = node
	} else {
		node.next = linkedList.head
		linkedList.head = node
	}

	linkedList.size += 1
}

func (linkedList *LinkedList) RemoveFirst() (bool, int) {
	node := linkedList.head
	if !linkedList.Empty() {
		linkedList.head = linkedList.head.next
		linkedList.size -= 1
		return true, node.value
	}

	return false, 0
}

func (linkedList *LinkedList) AddLast(value int) {
	node := &LinkedListNode{value: value}

	if linkedList.size == 0 {
		linkedList.head = node
		linkedList.tail = node
	} else {
		linkedList.tail.next = node
		linkedList.tail = node
	}

	linkedList.size += 1
}

func (linkedList *LinkedList) get(index int) (bool, *LinkedListNode) {
	currentIndex := 0
	node := linkedList.head

	if index >= 0 && index < linkedList.size {
		for node != nil {
			if currentIndex == index {
				return true, node
			}

			node = node.next
			currentIndex++
		}
	}

	return false, nil
}

func (linkedList *LinkedList) Get(index int) (bool, int) {
	status, node := linkedList.get(index)

	if status {
		return status, node.value
	}

	return false, 0
}

func (linkedList *LinkedList) Set(index int, value int) bool {
	status, node := linkedList.get(index)

	if status {
		node.value = value
	}

	return status
}

func (linkedList *LinkedList) ToString() string {
	var buffer bytes.Buffer
	node := linkedList.head

	for node != nil {
		buffer.WriteString(strconv.Itoa(node.value) + " ")
		node = node.next
	}
	return buffer.String()
}

func (linkedList *LinkedList) Head() (bool, int) {
	if !linkedList.Empty() {
		return true, linkedList.head.value
	}

	return false, 0
}

func (linkedList *LinkedList) Tail() (bool, int) {
	if !linkedList.Empty() {
		return true, linkedList.tail.value
	}

	return false, 0
}

func (linkedList *LinkedList) Empty() bool {
	return linkedList.size == 0
}

func (linkedList *LinkedList) Size() int {
	return linkedList.size
}

func (linkedList *LinkedList) Clear() {
	linkedList.head = nil
	linkedList.tail = nil
	linkedList.size = 0
}
