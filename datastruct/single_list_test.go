package datastruct

import "testing"

func TestInsert(t *testing.T) {
	// 初始化单链表实例
	t.Logf("TestInsert running ... \n")
	n := NewSingleList()
	// 从头部插入3个节点
	n.InsertData(1, 100)
	n.InsertData(1, 10)
	n.InsertData(1, 1)
	// 打印链表中所有节点数据
	n.Traverse()
	// 插入错误节点
	//n.InsertData(5,10)

	// 插入新节点
	node := new(SingleList)
	node.Data = 200
	n.InsertSingleList(2, node)
	n.Traverse()
}

func TestGetList(t *testing.T) {
	t.Logf("TestGetList running ... \n")
	// 初始化单链表实例
	n := NewSingleList()
	// 从头部插入4个节点
	n.InsertData(1, 1000)
	n.InsertData(1, 100)
	n.InsertData(1, 10)
	n.InsertData(1, 1)
	// 打印链表长度
	len := n.GetLen()
	t.Logf("length of the SingleList is: %d\n", len)
	// 获取第一个节点
	head := n.GetFirstSingleList()
	t.Logf("head of the SingleList is: %d\n", head.Data)
	// 获取最后一个节点
	tail := n.GetLastSingleList()
	t.Logf("tail of the SingleList is: %d\n", tail.Data)
	// 获取链表中第3个节点
	if node, b := n.GetSingleList(3); b == true {
		t.Logf("the node is: %d\n", node.Data)
	}else {
		t.Logf("The position of the node is wrong!\n")
	}
	// 获取链表中第2个节点数据
	if nodedata, b := n.GetData(2); b == true {
		t.Logf("the node data is: %d\n", nodedata)
	} else {
		t.Logf("The position of the node is wrong!\n")
	}
}

func TestDeleteList(t *testing.T) {
	t.Logf("TestDeleteList running ... \n")
	// 初始化单链表实例
	n := NewSingleList()
	// 从头部插入4个节点
	n.InsertData(1, 1000)
	n.InsertData(1, 100)
	n.InsertData(1, 10)
	n.InsertData(1, 1)
	t.Logf("before delete...")
	n.Traverse()
	n.Delete(1)
	t.Logf("after delete...")
	n.Traverse()
}