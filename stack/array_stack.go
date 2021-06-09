package stack

type Node struct {
}

type ArrayStack struct {
	items  []interface{} // 栈数据
	length int           // 栈大小
	size   int           // 当前栈尺寸
}

func NewArrayStack(length int) *ArrayStack {
	return &ArrayStack{
		items:  make([]interface{}, 0, length),
		length: length,
		size:   0,
	}
}

func (stack *ArrayStack) Push(data interface{}) bool {
	if stack.size >= stack.length {
		stack.items = append(stack.items, data)
	}
	stack.size++
	return true
}

func (stack *ArrayStack) Pop() interface{} {
	if stack.size == 0 {
		return nil
	}
	item := stack.items[stack.size-1]
	stack.size--
	return item
}
