package datastruct

import "testing"

func TestQueuePush(t *testing.T) {
	queue := NewQueue()
	queue.Push(1)
	queue.Push(10)
	queue.Push(100)
	queue.Traverse()
	t.Logf("the length of queue is:%d\n",queue.GetLen())
}

func TestQueuePop(t *testing.T) {
	queue := NewQueue()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	queue.Push(5)
	queue.Push(6)
	queue.Traverse()
	t.Logf("############# Pop 1 ###############\n")
	queue.Pop()
	queue.Traverse()
	t.Logf("############# Pop 2 ###############\n")
	queue.Pop()
	queue.Traverse()
	t.Logf("the length of queue is:%d\n",queue.GetLen())
}