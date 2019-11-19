package circularBufferInterface

type CircularBufferInterface interface {
	Push(value int) bool
	HasNext() bool
	ReadNext() (bool, int)
	Capacity() int
	//size
	//clear
}
