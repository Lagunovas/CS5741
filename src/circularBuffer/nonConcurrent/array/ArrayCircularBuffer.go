package arrayCircularBuffer

type ArrayCircularBuffer struct {
	data     []int
	read     int
	write    int
	size     int
	capacity int
}

func NewArayCircularBuffer(capacity int) *ArrayCircularBuffer {
	return &ArrayCircularBuffer{[]int{}, 0, 0, 0, capacity}
}

func (arrayCircularBuffer *ArrayCircularBuffer) Push(value int) bool {
	if arrayCircularBuffer.canWrite() {
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

func (arrayCircularBuffer *ArrayCircularBuffer) HasNext() bool {
	return arrayCircularBuffer.size > 0
}

func (arrayCircularBuffer *ArrayCircularBuffer) ReadNext() (bool, int) {
	var value int

	if arrayCircularBuffer.HasNext() {
		if arrayCircularBuffer.read == arrayCircularBuffer.capacity {
			arrayCircularBuffer.read = 0
		}

		value = arrayCircularBuffer.data[arrayCircularBuffer.read]
		arrayCircularBuffer.size--
		arrayCircularBuffer.read++
		return true, value
	}

	return false, value
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

func (arrayCircularBuffer *ArrayCircularBuffer) canWrite() bool {
	return arrayCircularBuffer.size < arrayCircularBuffer.capacity
}
