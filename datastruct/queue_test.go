package datastruct

import "testing"

func TestQueuePush(t *testing.T) {
	queue := NewQueue()
	queue.Push(1)
	queue.Push(10)
	queue.Push(100)
	queue.ForwardTraverse()
	t.Logf("the length of queue is:%d\n", queue.GetLen())
	queue.ReverseTraverse()
	t.Logf("the head of queue value: %v\n", queue.Peek())
}

func TestQueuePop(t *testing.T) {
	queue := NewQueue()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	queue.Push(5)
	queue.Push(6)
	queue.ForwardTraverse()
	t.Logf("############# Pop 1 ###############\n")
	data1, _ := queue.Pop()
	t.Logf("the Pop data: %v\n", data1)
	queue.ForwardTraverse()
	t.Logf("############# Pop 2 ###############\n")
	data2, _ := queue.Pop()
	t.Logf("the Pop data: %v\n", data2)
	t.Logf("the head of queue value: %v\n", queue.Peek())
	t.Logf("the length of queue is:%d\n", queue.GetLen())
}
