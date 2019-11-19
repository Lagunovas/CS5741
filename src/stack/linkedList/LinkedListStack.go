package linkedListStack

import (
	"fmt"
	linkedList "github.com/CS5741/src/shared/linkedList"
)

//=========>Stack<==============
type Stack struct{
	list *linkedList.LinkedList
}

func NewStack() *Stack{
	list := linkedList.NewLinkedList()
	return &Stack{list}
}

func (stack *Stack)Push(value int){
	stack.list.AddFirst(value)
}

/*============> MAIN<<==============*/
// func main() {
// 	fmt.Println("This is a test")
	
// 	stack  := NewStack()
// 	stack.Push(1)
// 	stack.Push(2)
// 	stack.Push(3)
// 	stack.peek()
// 	stack.Pop()
// 	stack.Push(4)
// 	stack.peek()
// }

func (stack *Stack) Pop() (bool, int) {
	var value int
	if !stack.Empty(){
		_ , val := stack.list.RemoveFirst() 
		fmt.Printf("Pop: %d \n", val)
		return true, val
	} else {
		fmt.Printf("the stack is empty \n")
		return false, value
	}
	
}

//peek 
func (stack *Stack) Peek() (bool, int) {
	var value int
	if !stack.Empty(){
		_ , val := stack.list.GetHead()
		fmt.Printf("peek: %d \n", val)
		return true, val
	}else{
		return false, value
	}
}
//isEmpty 
func (stack *Stack) Empty() bool{
	return stack.list.Empty()
}
//size 
func (stack *Stack) Size() int{
	return stack.list.Size()
}
//clear 
func (stack *Stack) Clear() {
	stack.list = linkedList.NewLinkedList()
}

