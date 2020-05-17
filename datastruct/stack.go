package datastruct

import "log"

type ElementData interface{}

type Stack struct {
	Top    *Node
	Length int // 栈的长度
}

type Node struct {
	value ElementData
	next  *Node
}

func (n *Node) Value() ElementData {
	return n.value
}

// 创建栈对象
func NewStack() *Stack {
	return &Stack{nil, 0}
}

// 返回栈的长度
func (this *Stack) Len() int {
	return this.Length
}

// 入栈
func (this *Stack) Push(v ElementData) {
	data := &Node{
		value: v,
		next:  this.Top,
	}
	this.Top = data
	this.Length += 1
}

// 出栈
func (this *Stack) Pop() ElementData {
	if this.Len() == 0 {
		return nil
	}
	result := this.Top.Value()
	this.Top = this.Top.next
	this.Length -= 1
	return result
}

// 查看栈顶元素
func (this *Stack) Peek() ElementData {
	if this.Len() == 0 {
		return nil
	}
	return this.Top.Value()
}

// 遍历栈, 打印数据
func (this *Stack) Traverse() {
	top := this.Top
	for top != nil {
		log.Printf("The Data of Stack is:%v\n", top.Value())
		top = top.next
	}
}
