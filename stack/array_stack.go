package stack

type Node struct {
}

type ArrayStack struct {
	Items  []interface{}
	Length int
	Size   int
}

func NewArrayStack(length int) *ArrayStack {
	return &ArrayStack{
		Items:  make([]interface{}, 0, length),
		Length: length,
		Size:   0,
	}
}
