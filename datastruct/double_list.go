package datastruct

import (
	"errors"
	"log"
)

// 链表数据类型
type Element interface{}

type LinkedList struct {
	Length int // 链表的长度
	List   *Double_list
}

type Double_list struct {
	prev *Double_list
	next *Double_list
	data Element
}

// error定义
var (
	ErrIndexOutOfBound = errors.New("the location out of index range")
	ErrEmptyList       = errors.New("the list is empty")
	ErrNotFoundNode    = errors.New("not found node in the list")
)

// 创建链表结构实例
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// 获取链表长度
func (ll *LinkedList) GetLength() int {
	return ll.Length
}

// 将元素组装成节点
func (ll *LinkedList) newElement(v ...Element) []*Double_list {
	elements := []*Double_list{}
	for index, value := range v {
		element := new(Double_list)
		element.data = value
		if index != 0 {
			element.prev = elements[index-1]
			element.prev.next = element
		}
		elements = append(elements, element)
	}
	return elements
}

// 查找指定位置，获取该位置节点
func (ll *LinkedList) findLocation(location int) (*Double_list, error) {
	// 判断位置是否有效
	if location < 0 || location > ll.Length {
		return nil, ErrIndexOutOfBound
	}
	if ll.Length == 0 {
		return nil, ErrEmptyList
	}
	entry := ll.List
	index := 0
	for {
		if index == location {
			return entry, nil
		}
		if entry.HasNext() {
			index++
			entry = entry.Next()
		} else {
			break
		}
	}
	return nil, ErrNotFoundNode
}

// 插入元素到指定位置
func (ll *LinkedList) Insert(location int, v ...Element) error {
	elements := ll.newElement(v...)
	// 链表为空
	if ll.Length == 0 {
		ll.List = elements[0]
		ll.Length = len(elements)
		return nil
	}

	// 插入尾部
	if location == ll.Length {
		if l, err := ll.GetLast(); err != nil {
			return err
		} else {
			l.next = elements[0]
			elements[0].prev = l
			ll.Length += len(elements)
			return nil
		}
	}

	// 链表非空
	// 查找指定位置
	entry, err := ll.findLocation(location)
	if err != nil {
		return err
	}

	elements[0].prev = entry.prev
	elements[len(elements)-1].next = entry

	if entry.prev != nil {
		entry.prev.next = elements[0]
	}

	if entry != nil {
		entry.prev = elements[len(elements)-1]
	}
	// 插入位置为头部，更新ll.list指向
	if location == 0 {
		ll.List = elements[0]
	}

	ll.Length += len(elements)

	return nil
}

// 插入元素到头部
func (ll *LinkedList) InsertFirst(v Element) error {
	return ll.Insert(0, v)
}

// 清空链表
func (ll *LinkedList) Delete() {
	ll.List = nil
	ll.Length = 0
}

// 检查链表中是否包含指定元素
func (ll *LinkedList) Contains(v Element) bool {
	contain := false
	ll.Iterator(func(ind int, node *Double_list) bool {
		if node.data == v {
			contain = true
			return true
		}
		return false
	})
	return contain
}

// 获取指定位置的节点
func (ll *LinkedList) Get(location int) (*Double_list, error) {
	if location < 0 || location >= ll.Length {
		return nil, ErrIndexOutOfBound
	}
	if ll.Length == 0 {
		return nil, ErrNotFoundNode
	}
	var result *Double_list
	ll.Iterator(func(ind int, node *Double_list) bool {
		if ind == location {
			result = node
			return true
		}
		return false
	})
	return result, nil
}

// 获取头部节点
func (ll *LinkedList) GetFirst() (*Double_list, error) {
	return ll.Get(0)
}

// 获取尾部节点
func (ll *LinkedList) GetLast() (*Double_list, error) {
	return ll.Get(ll.Length - 1)
}

// 获取元素所在的索引
func (ll *LinkedList) GetIndex(v Element) int {
	var index int
	ll.Iterator(func(ind int, node *Double_list) bool {
		if node.data == v {
			index = ind
			return true
		}
		return false
	})
	return index
}

// 遍历打印链表
func (ll *LinkedList) Traverse() {
	ll.Iterator(func(ind int, node *Double_list) bool {
		log.Printf("The Data of DoubleList is:%v\n", node.data)
		return false
	})
}

// 移除指定位置节点
func (ll *LinkedList) Remove(location int) error {
	entry, err := ll.findLocation(location)
	if err != nil {
		return err
	}

	if entry.prev != nil {
		entry.prev.next = entry.next
	}
	if entry.next != nil {
		entry.next.prev = entry.prev
	}
	// 当移除头部节点时，更新list的指向
	if location == 0 {
		ll.List = entry.next
	}

	ll.Length -= 1

	return nil
}

// 移除头部节点
func (ll *LinkedList) Pull() error {
	return ll.Remove(0)
}

// 移除尾部节点
func (ll *LinkedList) Pop() error {
	return ll.Remove(ll.Length - 1)
}

// 追加节点
func (ll *LinkedList) Push(v ...Element) error {
	return ll.Insert(ll.Length, v...)
}

// 修改指定位置元素内容
func (ll *LinkedList) Set(location int, v Element) error {
	ll.Iterator(func(ind int, node *Double_list) bool {
		if ind == location {
			node.data = v
			return true
		}
		return false
	})
	return nil
}

// 移动元素到指定位置
func (ll *LinkedList) Move(location int, list *Double_list) error {
	var (
		srclen = ll.Length
	)
	if location < 0 || location > srclen {
		return ErrIndexOutOfBound
	}
	entry, srcSecond := ll.List, ll.List.next
	index, listLoc := 0, 0
	var entryDest, entrySrc *Double_list
	destFlat, srcFlag := false, false
	for {
		if entry != nil {
			if index == location {
				entryDest = entry
				destFlat = true
			}
			if entry.data == list.data {
				listLoc = index
				entrySrc = entry
				srcFlag = true
			}
			if destFlat && srcFlag {
				break
			}
			if entry.HasNext() {
				index++
				entry = entry.Next()
			} else {
				break
			}
		}
	}

	// 指定位置的节点未找到，或待移动的元素未找到
	if (!destFlat && location != ll.Length) || !srcFlag {
		return ErrNotFoundNode
	}
	if location == listLoc {
		return nil
	}

	// 移动元素，更改该元素前后节点的指向
	if entrySrc.prev != nil {
		entrySrc.prev.next = entrySrc.next
	}
	if entrySrc.next != nil {
		entrySrc.next.prev = entrySrc.prev
	}
	ll.Length -= 1

	// 插入到尾部
	if location == srclen {
		if listLoc == 0 {
			ll.List = srcSecond
		}
		if l, err := ll.GetLast(); err != nil {
			return err
		} else {
			l.next = entrySrc
			entrySrc.prev = l
			entrySrc.next = nil
		}
	} else {
		// 插入到指定节点
		if listLoc < location { // 如果移动元素在目的位置前，插入到目的位置后
			entrySrc.prev = entryDest
			entrySrc.next = entryDest.next

			entryDest.next = entrySrc
			if entryDest.next != nil {
				entryDest.next.prev = entrySrc
			}

		} else { // 如果移动元素在目的位置后，插入到目的位置前
			entrySrc.prev = entryDest.prev
			entrySrc.next = entryDest

			if entryDest.prev != nil {
				entryDest.prev.next = entrySrc
			}
			if entryDest != nil {
				entryDest.prev = entrySrc
			}
		}

	}
	ll.Length += 1

	// 链表头部移动了
	if listLoc == 0 {
		ll.List = srcSecond
	}
	// 移动元素到链表头部
	if location == 0 {
		ll.List = entrySrc
	}

	return nil
}

// 移动元素到头部
func (ll *LinkedList) MoveFront(list *Double_list) error {
	return ll.Move(0, list)
}

// 移动元素到尾部
func (ll *LinkedList) MoveBack(list *Double_list) error {
	return ll.Move(ll.Length-1, list)
}

// 获取链表返回切片
func (ll *LinkedList) AsList() []Element {
	result := []Element{}
	ll.Iterator(func(ind int, node *Double_list) bool {
		result = append(result, node.data)
		return false
	})
	return result
}

// 遍历列表，遍历后的操作由参数传入
func (ll *LinkedList) Iterator(iterator func(ind int, node *Double_list) bool) {
	if ll.Length == 0 {
		return
	}
	index := 0
	list := ll.List

	for {
		result := iterator(index, list)
		if result {
			break
		}
		if list.HasNext() {
			index++
			list = list.Next()
		} else {
			break
		}
	}
	return
}

// 判断是否还有next节点
func (dl *Double_list) HasNext() bool {
	if dl.next != nil {
		return true
	}
	return false
}

// 获取下一个节点
func (dl *Double_list) Next() *Double_list {
	return dl.next
}

func (dl *Double_list) Value() Element {
	return dl.data
}
