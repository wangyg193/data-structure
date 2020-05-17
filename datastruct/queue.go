package datastruct

import (
	"errors"
	"log"
)

type ElementList interface{}

type Queue struct {
	head   *list
	tail   *list
	length int // 队列的长度
}

type list struct {
	value ElementList
	prev  *list
	next  *list
}

func (l *list) Value() ElementList {
	return l.value
}

// 创建队列结构实例
func NewQueue() *Queue {
	return &Queue{head: nil, tail: nil, length: 0}
}

// 查看队头元素
func (this *Queue) Peek() ElementObject {
	if this.head == nil {
		return nil
	}
	return this.head.Value()
}

// Push入队列，队尾插入
func (this *Queue) Push(data ElementObject) {
	element := &list{
		value: data,
		prev:  this.tail,
		next:  nil,
	}
	// 队列为空
	if this.tail == nil {
		this.head = element
		this.tail = element
	} else {
		this.tail.next = element
		this.tail = element
	}
	this.length += 1
}

// Pop出队列，队头移除
func (this *Queue) Pop() (ElementObject, error) {
	if this.head == nil {
		return nil, errors.New("Empty queue.")
	}
	head := this.head
	this.head = head.next
	head.next.prev = nil
	this.length--
	return head.Value(), nil
}

// 获取队列长度
func (this *Queue) GetLen() int {
	return this.length
}

// 从队头遍历队列, 打印数据
func (this *Queue) ForwardTraverse() {
	head := this.head
	for head != nil {
		log.Printf("ForwardTraverse: The Data of Queue is:%v\n", head.Value())
		head = head.next
	}
}

// 从队尾遍历队列，打印数据
func (this *Queue) ReverseTraverse() {
	tail := this.tail
	for tail != nil {
		log.Printf("ReverseTraverse: The Data of Queue is:%v\n", tail.Value())
		tail = tail.prev
	}
}