package linkedList

import (
	"fmt"
	"bytes"
	"strconv"
)

type LinkedListNode struct {
	value int 
	next *LinkedListNode	
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


//==========>traversing
func (list *LinkedList) ToString() string {
	var buffer bytes.Buffer
	node := list.head
	//fmt.Printf("the elements of the list are:\n")
	for node != nil{
		buffer.WriteString(strconv.Itoa(node.value) + " ")
		node = node.next
	}
	return buffer.String()
}  


//===========>addFirst
func (list *LinkedList) AddFirst(value int) {
	node := &LinkedListNode{value: value}
	if list.Empty() { //if the linked list has no elements in it 
		list.head = node
		list.tail = node
	} else {
		node.next = list.head
		list.head = node
	}
	list.size += 1
	//fmt.Printf("the node has been added at head\n")
}

//RemoveFirst
func (list *LinkedList) RemoveFirst() (bool, int) {
	var value int 

	node := list.head
	if !list.Empty() {
		list.head = list.head.next
		list.size -= 1
		//fmt.Printf("first element removed successfully\n")
		return true, node.value
	} else {
		//fmt.Printf("list is empty \n")
		return false, value
	}
	 
}

func (list *LinkedList) InsertAtIndex(value int, indexForInsertion  int) {
	index := 0
	node := list.head
	nodeToBeAdded := &LinkedListNode{value: value}
	if list.size == 0 {
		//fmt.Printf("the linked list is empty \n")
	} else if indexForInsertion >= list.size {
		//fmt.Printf("index entered exceeds the lenght of the linked list\n")
	} else {
		for index != indexForInsertion { 
			if index == indexForInsertion - 1 { // coz the one before this is needed
				temp := node.next
				node.next = nodeToBeAdded
				nodeToBeAdded.next = temp
				list.size += 1
				//fmt.Printf("inserted successfully before index %d \n", indexForInsertion)
			} else {
				node = node.next
			}
		 	index += 1
		} 
	}
}

func (list *LinkedList) GetNodeAtIndex(indexToBeReturned int) (bool, int){
	var value int 
	index := 0
	node := list.head

	if list.size == 0{
		//fmt.Printf("the linked list is empty false returned\n")
	}else if indexToBeReturned >= list.size{
		//fmt.Printf("index entered exceeds the lenght of the linked list nil returned\n")
	}else{
		if indexToBeReturned == 0{
			return list.GetHead()
		}
		if indexToBeReturned == list.size - 1{
			return list.GetTail()
		}
		for index != indexToBeReturned { 
			if index == indexToBeReturned - 1{ // coz the one before this is needed
				return true, node.next.value
				//fmt.Printf("inserted successfully \n")
			}else{
				node = node.next
			}
		 	index += 1
		 } 
	}
	return false, value
}

func (list *LinkedList) AddLast(value int){
	node := &LinkedListNode{value: value}
	
	if list.size == 0{
		list.head = node
		list.tail = node
	}else{
		list.tail.next = node
		list.tail = node 	
	}
	list.size += 1
	//fmt.Printf("the node has been successfully added at tail\n")
}

func (list *LinkedList) EditNodeAtIndex(newValue int, indexToBeEdited int) bool {
	index := 0
	node := list.head

	if list.size == 0{
		//fmt.Printf("the linked list is empty false returned\n")
	}else if indexToBeEdited >= list.size{
		fmt.Printf("index entered exceeds the lenght of the linked list nil returned\n")
	}else{
		if indexToBeEdited == 0{
			list.head.value = newValue
		}
		for index != indexToBeEdited + 1 { 
			if index == indexToBeEdited{ 
				node.value = newValue
				//fmt.Printf("node edited successfully \n")
				return true
			}else{
				node = node.next
			}
		 	index += 1
		 } 
	}
	return false
}

func (list *LinkedList) Empty() bool {
	return list.size <= 0
}

func (list *LinkedList) GetHead() (bool, int) {
	var value int 
	if !list.Empty() {
		return true, list.head.value
	}
	return false, value
}

func (list *LinkedList) Size() int {
	return list.size
}

func (list *LinkedList) GetTail() (bool, int) {
	var value int
	if !list.Empty(){
		return true, list.tail.value
	}
	return false, value
}