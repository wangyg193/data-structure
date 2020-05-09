package datastruct

import "testing"

func TestQueuePush(t *testing.T) {
	queue := NewQueue()
	queue.Push(1)
	queue.Push(10)
	queue.Push(100)
	queue.Traverse()
}