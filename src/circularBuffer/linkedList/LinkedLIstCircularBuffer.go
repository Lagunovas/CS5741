package linkedListCircularBuffer

//package main


import(
	linkedList "github.com/CS5741/src/shared/linkedList"
	"fmt"

)
//======================>Circular Buffer<====================


type CircularBuffer struct{
	list *linkedList.LinkedList
	read int
	write int
	size int 
	capacity int
}


func NewCircularBuffer(bufferCapacity int) *CircularBuffer{
	list := linkedList.NewLinkedList()
	return &CircularBuffer{list, 0, 0,0,bufferCapacity}
}

/*============> MAIN<<==============*/
func main() {
	fmt.Println("This is a test")
	buffer := NewCircularBuffer(5)
	buffer.ReadNext()
	buffer.Push(10)
	buffer.Push(20)
	buffer.Push(30)
	buffer.Push(40)
	buffer.Push(50)
	buffer.Push(60)
	fmt.Println(buffer.ToString())
	buffer.ReadNext()
	buffer.ReadNext()
	buffer.ReadNext()
	buffer.ReadNext()
	buffer.ReadNext()
	buffer.ReadNext()
	fmt.Println(buffer.ToString())

	// fmt.Println(buffer.ReadNext())
	// fmt.Println(buffer.ReadNext())
	// fmt.Println(buffer.ReadNext())
	// fmt.Println(buffer.ReadNext())
	// fmt.Println(buffer.ReadNext())
	// fmt.Println(buffer.ReadNext())
	// fmt.Println(buffer.ToString())
	buffer.Push(70)
	buffer.Push(80)
	buffer.Push(90)
	buffer.Push(100)
	fmt.Println(buffer.ToString())
	buffer.ReadNext()
	buffer.ReadNext()
	buffer.ReadNext()
	// fmt.Println(buffer.ReadNext())
	// fmt.Println(buffer.ReadNext())
	// fmt.Println(buffer.ReadNext())
	fmt.Println(buffer.ToString())
}

//===========>Push<================
func (buffer *CircularBuffer) Push(value int) bool{
	//size == capacity=> no space so no push possible
	if buffer.size < buffer.capacity  {		
		if buffer.write == buffer.capacity {
			//change the write to zero and push
			buffer.write = 0
		}
		status, _ := buffer.list.GetNodeAtIndex(buffer.write)
		if status{
			//node exists change the value of the node insert at index wont work for this 
			buffer.list.EditNodeAtIndex(value, buffer.write)
		}else{
			//Insert at end 
			buffer.list.AddLast(value)
		}
		
		buffer.size++
		buffer.write++
		//fmt.Printf("added successfully %d \n" , buffer.write)
		fmt.Printf("size %d \n " , buffer.size)
		fmt.Printf("write %d \n " , buffer.write)
		fmt.Printf("read %d \n " , buffer.read)
		return true
	}else{
		fmt.Println("push not possible")
		return false
	}
}

func (buffer *CircularBuffer) ReadNext() (bool, int){
	var value int 
	if buffer.HasNext(){
		_ , val :=	buffer.list.GetNodeAtIndex(buffer.read)
		fmt.Printf("read successful %d \n", buffer.read)
		if buffer.read == buffer.capacity - 1 {
			buffer.read = 0 
		} else {
			buffer.read++
		}
		buffer.size--
		if buffer.size < 0 {
			buffer.size = 0
		}
		fmt.Printf("size %d \n " , buffer.size)
		fmt.Printf("write %d \n " , buffer.write)
		fmt.Printf("read %d \n " , buffer.read)
		return true, val
	}
	fmt.Println("does not have anything to read")
	return false, value
}

func (buffer *CircularBuffer)ToString() string {
	return buffer.list.ToString()
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