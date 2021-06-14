package queue

import "testing"

func TestArrayQueue_Enqueue(t *testing.T) {
	queue := NewArrayQueue(5)
	queue.Enqueue("hello")
	queue.Enqueue("world")

	t.Log(queue.String())

	item := queue.Dequeue()
	t.Log("出队：", item)
	t.Log("队列：", queue.String())
	queue.Enqueue(123)
	queue.Enqueue(345)
	item = queue.Dequeue()
	t.Log("出队：", item)
	t.Log("队列：", queue.String())
	queue.Enqueue(567)
	t.Log("队列：", queue.String())
}
