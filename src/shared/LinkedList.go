package main

import (
	"fmt"
)

/*
	Linked List=> this is a singly linked list with following functionalities 
	1. AddFirst 
	2. AddLast 
	3. Traversal => to print the linked list
	4. InsertAfterIndex //index starts from 0  
	5. InsertBeforeIndex //index starts from 0  
	6. RemoveFirst
	7. RemoveLast
	8. RemoveAtIndex 
	9. Search 
	10. Copy => this will create a clone of the linked list 
	11. GetNodeAtIndex //index starts from 0 
	12. IsEmPTY
*/	
	/*
	=============================>	USAGE <======================
	linkList := NewLinkedList()
	linkList.RemoveFirst()
	linkList.InsertAfterIndex(5,1)
	linkList.InsertBeforeIndex(6,1)
	linkList.AddFirst(2)
	linkList.AddLast(3)
	linkList.AddFirst(4)
	linkList.InsertAfterIndex(5,1)
	linkList.InsertBeforeIndex(6,1)
	linkList.RemoveFirst()
	linkList.RemoveLast()
	linkList.AddLast(7)
	linkList.AddLast(8)
	linkList.RemoveAtIndex(1)
	linkList.Traversing()
	stack  := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.peek()
	stack.Pop()
	stack.Push(4)
	stack.peek()

	*/
type Node struct{
	data int 
	next *Node	
} 

func NewNode(data int) *Node{
	return &Node{data: data}
}

type LinkedList struct {
 	head *Node
 	tail *Node 
 	size int
}

func NewLinkedList() *LinkedList{
	return &LinkedList{nil, nil, 0}
}

/*============> MAIN<<==============*/
func main() {
	fmt.Println("This is a test")
	

	
}


 //===========>addFirst
func (list *LinkedList) AddFirst(data int){
	node := &Node{data: data}
	if list.size == 0{ //if the linked list has no elements in it 
		list.head = node
		list.tail = node
	}else{
		node.next = list.head
		list.head = node
	}
	list.size += 1
	fmt.Printf("the node has been added at head\n")
}

 //===========>addLast
func (list *LinkedList) AddLast(data int){
	node := &Node{data: data}
	
	if list.size == 0{
		list.head = node
		list.tail = node
	}else{
		list.tail.next = node
		list.tail = node 	
	}
	list.size += 1
	fmt.Printf("the node has been successfully added at tail\n")
}

 //==========>traversing
func (list *LinkedList) Traversing(){
	node := list.head
	fmt.Printf("the elements of the list are:\n")
	for node != nil{
		fmt.Printf("%d \n", node.data)
		node = node.next
	}
}  

/*cases of 0th and last index are not included */
 //Inserting After Index
//indexForInsertion is the index at which the new element will be added 
//index starts at 0 
func (list *LinkedList) InsertAfterIndex(data int, indexForInsertion  int){
	index := 0
	node := list.head
	nodeToBeAdded := &Node{data: data}
	if list.size == 0{
		fmt.Printf("the linked list is empty \n")
	}else if indexForInsertion >= list.size{
		fmt.Printf("index entered exceeds the lenght of the linked list\n")
	}else{
		for index != indexForInsertion + 1{ // the +1 will allow the if index == IndexForInsertion execute as without +1 this loop will never run for the required index
			if index == indexForInsertion{
				temp := node.next
				node.next = nodeToBeAdded
				nodeToBeAdded.next = temp
				list.size += 1
				fmt.Printf("inserted successfully after index %d \n", indexForInsertion)
			}else{
				node = node.next
			}
		 	index += 1
		 } 
	}
}

 //Inserting Before Index works same as Insert at index
func (list *LinkedList) InsertBeforeIndex(data int, indexForInsertion  int){
	index := 0
	node := list.head
	nodeToBeAdded := &Node{data: data}
	if list.size == 0{
		fmt.Printf("the linked list is empty \n")
	}else if indexForInsertion >= list.size{
		fmt.Printf("index entered exceeds the lenght of the linked list\n")
	}else{
		for index != indexForInsertion { 
			if index == indexForInsertion - 1{ // coz the one before this is needed
				temp := node.next
				node.next = nodeToBeAdded
				nodeToBeAdded.next = temp
				list.size += 1
				fmt.Printf("inserted successfully before index %d \n", indexForInsertion)
			}else{
				node = node.next
			}
		 	index += 1
		 } 
	}
}
// //Insert At Index
// func (list *LinkedList) InsertBeforeIndex(data int, indexForInsertion  int){
// 	index := 0
// 	node := list.head
// 	nodeToBeAdded := &Node{data: data}
// 	if list.size == 0{
// 		fmt.Printf("the linked list is empty \n")
// 	}else if indexForInsertion >= list.size{
// 		fmt.Printf("index entered exceeds the lenght of the linked list\n")
// 	}else{
// 		for index != indexForInsertion { 
// 			if index == indexForInsertion - 1{ // coz the one before this is needed
// 				temp := node.next
// 				node.next = nodeToBeAdded
// 				nodeToBeAdded.next = temp
// 				list.size += 1
// 				fmt.Printf("inserted successfully before index %d \n", indexForInsertion)
// 			}else{
// 				node = node.next
// 			}
// 		 	index += 1
// 		 } 
// 	}
// }
//RemoveFirst
func (list *LinkedList) RemoveFirst() *Node{
	node := list.head
	if list.size > 0{
		list.head = list.head.next
		list.size -= 1
		fmt.Printf("first element removed successfully\n")
		return node
	}else{
		fmt.Printf("list is empty \n")
		return nil
	}
	 
}
//RemoveLast
func (list *LinkedList) RemoveLast() *Node{
	node := list.head
	nodeToBeRemoved := list.tail
	if list.size > 0{
		if list.size == 1{
			list.head = nil
		}
		for node.next.next != nil{
			node = node.next
		}
		node.next = nil
		list.tail = node
		list.size -= 1
		return nodeToBeRemoved
	}else{
		fmt.Printf("list is empty \n")
	}
	fmt.Printf("removed tail \n")
	return nil
}

//RemoveAtIndex 
func (list *LinkedList) RemoveAtIndex(indexForRemoval int){
	index := 0
	node := list.head
	if list.size == 0{
		fmt.Printf("the linked list is empty \n")
	}else if indexForRemoval >= list.size{
		fmt.Printf("index entered exceeds the lenght of the linked list\n")
	}else{
		for index != indexForRemoval { 
			fmt.Printf("index %d", index)
			fmt.Printf("indexForRemoval %d", indexForRemoval)
			if index == indexForRemoval - 1{ // coz the one before this is needed
				next := node.next
				temp := next.next
				node.next = temp
				list.size -= 1
				fmt.Printf("removed successfully at index %d \n", indexForRemoval)
			}else{
				node = node.next
			}
		 	index += 1
		} 
	}
}



 search  
 //copy 

// getNodeAtIndex
func (list *LinkedList) GetNodeAtIndex(indexToBeReturned int) *Node{
	index := 0
	node := list.head

	if list.size == 0{
		fmt.Printf("the linked list is empty nil returned\n")
	}else if indexToBeReturned >= list.size{
		fmt.Printf("index entered exceeds the lenght of the linked list nil returned\n")
	}else{
		if indexToBeReturned == 0{
			return list.head
		}
		if indexToBeReturned == list.size - 1{
			return list.tail
		}
		for index != indexToBeReturned { 
			if index == indexToBeReturned - 1{ // coz the one before this is needed
				return node.next
				fmt.Printf("inserted successfully \n")
			}else{
				node = node.next
			}
		 	index += 1
		 } 
	}
	return nil
}

func (list *LinkedList) IsEmpty() bool{
	return list.size <= 0
}

func (list *LinkedList) GetHead() *Node{
	return list.head
}

func (list *LinkedList) GetTail() *Node{
	return list.tail
}

func (list *LinkedList) GetSize() int{
	return list.size
}


//=========>Stack<==============
type Stack struct{
	list *LinkedList
}

func NewStack() *Stack{
	list := NewLinkedList()
	return &Stack{list}
}

func (stack *Stack)Push(data int){
	stack.list.AddFirst(data)
}
func (stack *Stack)Pop() *Node {
	if stack.list.size > 0{
		node := stack.list.RemoveFirst() 
		stack.list.size -= 1
		fmt.Printf("Pop: %d \n", node.data)
		return node
	}else{
		fmt.Printf("the stack is empty \n")
		return nil
	}
	
}

//peek 
func (stack *Stack) peek() *Node{
	if stack.list.size > 0{
		fmt.Printf("peek: %d \n", stack.list.head.data)
		return stack.list.head
	}else{
		return nil
	}
}
//isEmpty 
func (stack *Stack) IsEmpty() bool{
	return stack.list.size == 0
}
//size 
func (stack *Stack) size() int{
	return stack.list.size
}
//clear 
func (stack *Stack) clear() {
	stack.list = NewLinkedList()
}

//======================>Circular Buffer<====================


type CircularBuffer struct{
	list *LinkedList
	read int
	write int
	size int 
	capacity int
}


func NewCircularBuffer(bufferCapacity int) *CircularBuffer{
	list := NewLinkedList()
	return &CircularBuffer{list, 0, 0,0,bufferCapacity}
}

//===========>Push<================
func (buffer *CircularBuffer) Push(data int) bool{//size == capacity=> no space so no push possible
	if buffer.size < buffer.capacity {
		node := NewNode(data)
		
		if buffer.write == buffer.capacity{
			//change the write to zero and push
			buffer.write = 0
		}
		buffer.list.InsertBeforeIndex(write)
		buffer.size++
		buffer.write++
		return true
	}else{
		//push not possible
		return false
	}
}

func (buffer *CircularBuffer) ReadNext() *Node{
	if buffer.HasNext(){
		node :=	buffer.list.getNodeAtIndex(buffer.read)
		buffer.read++
		buffer.size--
		return node
	}
	return nil
}

func (buffer *CircularBuffer) HasNext() bool{
	return buffer.size > 0

}

func (buffer *CircularBuffer) Capacity() int{
	return buffer.capacity
}

func (buffer *CircularBuffer) Size() int{
	return buffer.size
}