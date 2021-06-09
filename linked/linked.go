package linked_lists

// Node 单链表节点
type Node struct {
	Data interface{}
	Next *Node
}

// DoublyLinkedListNode 双向链表节点
type DoublyLinkedListNode struct {
	Data interface{}
	Prev *Node
	Next *Node
}
