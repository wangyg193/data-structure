/********************************************************************
// 功能：单链表结构实现，实现增、删、改。、查 操作
// 备注：第一个节点为头节点，并不真实保存数据
// 日期：2020-04-30
*********************************************************************/

package datastruct

import "log"

// 链表数据类型
type ElementStruct interface{}

type SingleList struct {
	Data ElementStruct
	Next *SingleList
}

// 创建链表结构实例
func NewSingleList() *SingleList {
	return &SingleList{Data: nil, Next: nil}
}

// 返回第一个节点
func (n *SingleList) GetFirstSingleList() *SingleList {
	if n.Next == nil {
		// 只有头节点，第一个节点不存在
		return nil
	}
	return n.Next
}

// 返回最后一个节点
func (n *SingleList) GetLastSingleList() *SingleList {
	if n.Next == nil {
		// 尾节点不存在
		return nil
	}

	h := n
	for h.Next != nil {
		h = h.Next
	}
	return h
}

// 在链表的第i个位置前插入一个元素，复杂度为o(n)
func (n *SingleList) InsertData(i int, d ElementStruct) bool {
	h := n
	j := 1
	// 查找第i-1节点
	for h != nil && j < i {
		h = h.Next
		j++
	}
	// 节点错误
	if h == nil || j > i {
		log.Fatalf("insert position:[%d] error\n", i)
		return false
	}

	// 添加节点
	node := &SingleList{Data: d}
	node.Next = h.Next
	h.Next = node

	return true
}

// 在链表的第i个位置前插入一个节点
func (n *SingleList) InsertSingleList(i int, e *SingleList) bool {
	h := n
	j := 1
	// 查找第i-1节点
	for h != nil && j < i {
		h = h.Next
		j++
	}
	// 节点错误
	if h == nil || j > i {
		log.Fatalf("insert position:[%d] error", i)
		return false
	}

	// 添加节点
	e.Next = h.Next
	h.Next = e

	return true
}

// 遍历链表, 打印数据
func (n *SingleList) Traverse() {
	h := n
	for h.Next != nil {
		h = h.Next
		log.Printf("The Data of SingleList is:%v\n", h.Data)
	}
}

// 获取链表中的第i个元素，复杂度为o(n)
func (n *SingleList) GetData(i int) (ElementStruct, bool) {
	// 位置有效性判断
	if i < 0 || i > n.GetLen() {
		return -1, false
	}

	h := n
	j := 0
	for h.Next != nil {
		h = h.Next
		j++ // 记录查找位置
		if j == i {
			// 查找到位置i，返回
			return h.Data, true
		}
	}
	return -1, false
}

// 获取链表中的第i个节点，复杂度为o(n)
func (n *SingleList) GetSingleList(i int) (*SingleList, bool) {
	if i < 0 || i > n.GetLen() {
		return nil, false
	}

	h := n
	j := 0
	for h.Next != nil {
		h = h.Next
		j++
		if j == i {
			return h, true
		}
	}
	return nil, false
}

// 获取链表中节点的长度
func (n *SingleList) GetLen() int {
	h := n
	j := 0
	for h.Next != nil {
		h = h.Next
		j++
	}
	return j
}

// 删除链表中第i个节点
func (n *SingleList) Delete(i int) bool {
	h := n
	j := 1
	// 查找第i-1个节点
	for h != nil && j < i {
		j++
		h = h.Next
	}

	if h == nil || j > i {
		log.Fatalf("Failed to delete the [%d]-th node\n", i)
		return false
	}

	h.Next = h.Next.Next
	return true
}

