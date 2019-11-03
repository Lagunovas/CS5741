// package main
package arrayCircularBuffer

type ArrayCircularBuffer struct {
	data     []int
	read     int
	write    int
	size     int
	capacity int
}

func NewArayCircularBuffer(initialCapacity int) *ArrayCircularBuffer {
	return &ArrayCircularBuffer{[]int{}, 0, 0, 0, initialCapacity}
}

func (arrayCircularBuffer *ArrayCircularBuffer) Push(value int) bool {
	if arrayCircularBuffer.size < arrayCircularBuffer.capacity {
		if arrayCircularBuffer.write == arrayCircularBuffer.capacity {
			arrayCircularBuffer.write = 0
		}

		arrayCircularBuffer.data = append(arrayCircularBuffer.data, value)
		arrayCircularBuffer.data[arrayCircularBuffer.write] = value
		arrayCircularBuffer.write++
		arrayCircularBuffer.size++
		return true
	}

	return false
}

func (arrayCircularBuffer *ArrayCircularBuffer) ReadNext() int {
	var value int

	if arrayCircularBuffer.HasNext() {
		if arrayCircularBuffer.read == arrayCircularBuffer.capacity {
			arrayCircularBuffer.read = 0
		}

		value = arrayCircularBuffer.data[arrayCircularBuffer.read]
		arrayCircularBuffer.size--
		arrayCircularBuffer.read++
		return value
	}

	return value
}

func (arrayCircularBuffer *ArrayCircularBuffer) HasNext() bool {
	return arrayCircularBuffer.size > 0
}

func (arrayCircularBuffer *ArrayCircularBuffer) Capacity() int {
	return arrayCircularBuffer.capacity
}

func (arrayCircularBuffer *ArrayCircularBuffer) Size() int {
	return arrayCircularBuffer.size
}

func (arrayCircularBuffer *ArrayCircularBuffer) Clear() {
	arrayCircularBuffer.size = 0
	arrayCircularBuffer.write = 0
	arrayCircularBuffer.read = 0
	arrayCircularBuffer.data = []int{}
}

// func main() {
// 	var circularBuffer = NewCircularBuffer(5)
// 	circularBuffer.Push(1)
// 	fmt.Printf("Size: %d\n", circularBuffer.Size())
// 	circularBuffer.Push(2)
// 	fmt.Printf("Size: %d\n", circularBuffer.Size())
// 	circularBuffer.Push(3)
// 	circularBuffer.Push(4)
// 	circularBuffer.Push(5)
// 	// circularBuffer.Push(3)
// 	// circularBuffer.Push(3)
// 	// circularBuffer.Push(3)
// 	// circularBuffer.Push(3)
// 	fmt.Printf("Size: %d\n", circularBuffer.Size())
// 	fmt.Printf("Capacity: %d\n", circularBuffer.Capacity())

// 	fmt.Printf("%v\n", circularBuffer.ReadNext())
// 	fmt.Printf("%v\n", circularBuffer.ReadNext())
// 	fmt.Printf("%v\n", circularBuffer.ReadNext())
// 	fmt.Printf("%v\n", circularBuffer.ReadNext())
// 	fmt.Printf("%v\n", circularBuffer.ReadNext())
// 	fmt.Printf("%v\n", circularBuffer.ReadNext()) // nil

// 	circularBuffer.Push(6)
// 	circularBuffer.Push(7)
// 	circularBuffer.Push(8)

// 	fmt.Printf("%v\n", circularBuffer.ReadNext())
// 	fmt.Printf("%v\n", circularBuffer.ReadNext())
// 	fmt.Printf("%v\n", circularBuffer.ReadNext())

// 	fmt.Printf("%v\n", circularBuffer.Push(9))
// 	fmt.Printf("%v\n", circularBuffer.Push(10))
// 	fmt.Printf("%v\n", circularBuffer.Push(11))
// 	fmt.Printf("%v\n", circularBuffer.Push(12))
// 	fmt.Printf("%v\n", circularBuffer.Push(13))
// 	fmt.Printf("%v\n", circularBuffer.Push(14))

// 	// fmt.Printf("%v\n", circularBuffer.ReadNext())
// }
