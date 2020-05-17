package datastruct

import "testing"

func TestStackPush(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	t.Logf("the top of stack value: %v\n",stack.Peek())
	stack.Traverse()
}

func TestStackPop(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Traverse()
	t.Logf("############# Pop 1 ###############\n")
	stack.Pop()
	stack.Traverse()
	t.Logf("############# Pop 2 ###############\n")
	stack.Pop()
	stack.Traverse()
}
