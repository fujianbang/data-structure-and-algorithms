package queue

import "encoding/json"

type arrayQueue struct {
	items    []interface{} // 队列中数据
	capacity int           // 容量
	head     int           // 队列头下标
	tail     int           // 队列尾下标
}

func NewArrayQueue(capacity int) *arrayQueue {
	return &arrayQueue{
		items:    make([]interface{}, capacity),
		capacity: capacity,
		head:     0,
		tail:     0,
	}
}

// Enqueue 入队
func (queue *arrayQueue) Enqueue(data interface{}) bool {
	if queue.tail >= queue.capacity {
		// 队列已满
		return false
	}
	queue.items[queue.tail] = data
	queue.tail++
	return true
}

// Dequeue 出队
func (queue *arrayQueue) Dequeue() interface{} {
	if queue.head >= queue.tail {
		// 队列空
		return false
	}
	item := queue.items[queue.head]
	queue.head++
	return item
}

// Capacity 队列容量
func (queue *arrayQueue) Capacity() int {
	return queue.capacity
}

// Size 队列中数据数量
func (queue *arrayQueue) Size() int {
	return queue.tail - queue.head
}

// String 格式化输出
func (queue *arrayQueue) String() string {
	list := queue.items[queue.head:queue.tail]
	bytes, _ := json.Marshal(list)
	return string(bytes)
}
