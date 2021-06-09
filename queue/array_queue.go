package queue

type ArrayQueue struct {
	items    []interface{} // 队列中数据
	capacity int           // 容量
	head     int           // 队列头下标
	tail     int           // 队列尾下标
}

func NewArrayQueue(capacity int) *ArrayQueue {
	return &ArrayQueue{
		items:    make([]interface{}, capacity),
		capacity: capacity,
		head:     0,
		tail:     0,
	}
}

func (queue *ArrayQueue) Enqueue(data interface{}) bool {
	if queue.tail >= queue.capacity {
		// 队列已满
		return false
	}
	queue.items[queue.tail] = data
	queue.tail++
	return true
}

func (queue *ArrayQueue) Dequeue() interface{} {
	if queue.head >= queue.tail {
		// 队列空
		return false
	}
	item := queue.items[queue.head]
	queue.head++
	return item
}
