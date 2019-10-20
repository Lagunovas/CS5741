package stack

type ArrayStack struct {
	items []interface{}
	size  int
}

func NewStack() *ArrayStack {
	return &ArrayStack{
		items: []interface{}{},
		size:  0,
	}
}

func (arrayStack *ArrayStack) Push(value interface{}) {
	arrayStack.items = append(arrayStack.items, value)
	arrayStack.size++
}

func (arrayStack *ArrayStack) Pop() interface{} {
	var item interface{}

	if !arrayStack.Empty() {
		item = arrayStack.items[arrayStack.size-1]
		arrayStack.items = arrayStack.items[0 : arrayStack.size-1]
		arrayStack.size--
	}

	return item
}

func (arrayStack *ArrayStack) Peek() *interface{} {
	var item *interface{}

	if !arrayStack.Empty() {
		item = &arrayStack.items[arrayStack.size-1]
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
