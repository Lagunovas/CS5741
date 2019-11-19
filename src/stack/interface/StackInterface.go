package stackInterface

type StackInterface interface {
	Push(value int)
	Pop() (bool, int)
	Peek() (bool, int)
	Empty() bool
	Size() int
	Clear()
}
