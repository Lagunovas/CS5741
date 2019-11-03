package main

import "fmt"

type CircularBuffer struct {
	data     []interface{}
	read     int
	write    int
	size     int
	capacity int
}

func NewCircularBuffer(initialCapacity int) *CircularBuffer {
	return &CircularBuffer{[]interface{}{}, 0, 0, 0, initialCapacity}
}

func (circularBuffer *CircularBuffer) Push(value interface{}) bool {
	if circularBuffer.size < circularBuffer.capacity {
		if circularBuffer.write == circularBuffer.capacity {
			circularBuffer.write = 0
		}

		circularBuffer.data = append(circularBuffer.data, value)
		circularBuffer.data[circularBuffer.write] = value
		circularBuffer.write++
		circularBuffer.size++
		return true
	}

	return false
}

// Removes oldest element (https://en.wikipedia.org/wiki/Circular_buffer)
func (circularBuffer *CircularBuffer) Pop() interface{} {
	return nil
}

func (circularBuffer *CircularBuffer) ReadNext() interface{} {
	var value interface{} = nil

	if circularBuffer.HasNext() {
		if circularBuffer.read == (circularBuffer.capacity - 1) {
			circularBuffer.read = 0
			fmt.Print("reset read pointer")
		}

		value = circularBuffer.data[circularBuffer.read]
		circularBuffer.size--
		circularBuffer.read++
	}

	return value
}

func (circularBuffer *CircularBuffer) HasNext() bool {
	return circularBuffer.size > 0
}

func (circularBuffer *CircularBuffer) Capacity() int {
	return circularBuffer.capacity
}

func (circularBuffer *CircularBuffer) Size() int {
	return circularBuffer.size
}

func main() {
	var circularBuffer = NewCircularBuffer(5)
	circularBuffer.Push(1)
	fmt.Printf("Size: %d\n", circularBuffer.Size())
	circularBuffer.Push(2)
	fmt.Printf("Size: %d\n", circularBuffer.Size())
	circularBuffer.Push(3)
	circularBuffer.Push(3)
	circularBuffer.Push(3)
	circularBuffer.Push(3)
	circularBuffer.Push(3)
	circularBuffer.Push(3)
	circularBuffer.Push(3)
	fmt.Printf("Size: %d\n", circularBuffer.Size())
	fmt.Printf("Capacity: %d\n", circularBuffer.Capacity())

	fmt.Printf("%v\n", circularBuffer.ReadNext())
	fmt.Printf("%v\n", circularBuffer.ReadNext())
	fmt.Printf("%v\n", circularBuffer.ReadNext())
	fmt.Printf("%v\n", circularBuffer.ReadNext())
}
