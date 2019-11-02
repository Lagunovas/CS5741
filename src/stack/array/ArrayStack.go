package arrayStack

type ArrayStack struct {
	items []int
	size  int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{make([]int, 0), 0}
}

func (arrayStack *ArrayStack) Push(value int) {
	arrayStack.items = append(arrayStack.items, value)
	arrayStack.size++
}

func (arrayStack *ArrayStack) Pop() int {
	var item int

	if !arrayStack.Empty() {
		item = arrayStack.items[arrayStack.size-1]
		arrayStack.items = arrayStack.items[0 : arrayStack.size-1]
		arrayStack.size--
	}

	return item
}

func (arrayStack *ArrayStack) Peek() int {
	var item int

	if !arrayStack.Empty() {
		item = arrayStack.items[arrayStack.size-1]
	}

	return item
}

func (arrayStack *ArrayStack) Empty() bool {
	return arrayStack.size == 0
}

func (arrayStack *ArrayStack) Size() int {
	return arrayStack.size
}

func (arrayStack *ArrayStack) Clear() {
	arrayStack.items = nil
	arrayStack.size = 0
}
