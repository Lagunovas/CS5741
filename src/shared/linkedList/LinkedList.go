package linkedList

import (
	"bytes"
	"fmt"
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

func (linkedList *LinkedList) GetNodeAtIndex(indexToBeReturned int) (bool, int) {
	var value int
	index := 0
	node := linkedList.head

	if linkedList.size == 0 {
		//fmt.Printf("the linked list is empty false returned\n")
	} else if indexToBeReturned >= linkedList.size {
		//fmt.Printf("index entered exceeds the lenght of the linked list nil returned\n")
	} else {
		if indexToBeReturned == 0 {
			return linkedList.Head()
		}
		if indexToBeReturned == linkedList.size-1 {
			return linkedList.Tail()
		}
		for index != indexToBeReturned {
			if index == indexToBeReturned-1 { // coz the one before this is needed
				return true, node.next.value
				//fmt.Printf("inserted successfully \n")
			} else {
				node = node.next
			}
			index += 1
		}
	}

	return false, value
}

func (linkedList *LinkedList) EditNodeAtIndex(newValue int, indexToBeEdited int) bool {
	index := 0
	node := linkedList.head

	if linkedList.size == 0 {
		//fmt.Printf("the linked list is empty false returned\n")
	} else if indexToBeEdited >= linkedList.size {
		fmt.Printf("index entered exceeds the lenght of the linked list nil returned\n")
	} else {
		if indexToBeEdited == 0 {
			linkedList.head.value = newValue
		}
		for index != indexToBeEdited+1 {
			if index == indexToBeEdited {
				node.value = newValue
				//fmt.Printf("node edited successfully \n")
				return true
			} else {
				node = node.next
			}
			index += 1
		}
	}
	return false
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
