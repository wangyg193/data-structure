package datastruct

import "testing"

func TestDoubleListInsert(t *testing.T) {
	list := NewLinkedList()
	elements := []Element{"hello", "world", "nihao", "shijie"}
	if err := list.Insert(0, elements...); err != nil {
		t.Fatalf("insert error:%v\n", err)
	}
	var e Element = "hi"
	if err := list.Insert(0, e); err != nil {
		t.Fatalf("insert error:%v\n", err)
	}
	var last Element = "bye"
	if err := list.Insert(5, last); err != nil {
		t.Fatalf("insert error:%v\n", err)
	}
	var mid []Element = []Element{"middle", "zhongjian"}
	if err := list.Insert(2, mid...); err != nil {
		t.Fatalf("insert error:%v\n", err)
	}

	head := []Element{"header", "testing"}
	if err := list.Insert(0, head...); err != nil {
		t.Fatalf("insert error:%v\n", err)
	}

	tail := []Element{"tail", "testing"}
	if err := list.Insert(list.Length, tail...); err != nil {
		t.Fatalf("insert error:%v\n", err)
	}

	if err := list.InsertFirst(1); err != nil {
		t.Fatalf("InsertFirst error:%v\n", err)
	}

	list.Traverse()
	return
}

func TestGetDoubleList(t *testing.T) {
	list := NewLinkedList()
	elements := []Element{"head", "hello", "world", "nihao", "shijie", "tail"}
	if err := list.Insert(0, elements...); err != nil {
		t.Fatalf("insert error:%v\n", err)
	}
	list.Traverse()

	if head, err := list.GetFirst(); err != nil {
		t.Fatalf("GetFirst error:%v\n", err)
	} else {
		t.Logf("head is %v\n", *head)
	}

	if tail, err := list.GetLast(); err != nil {
		t.Fatalf("GetLast error:%v\n", err)
	} else {
		t.Logf("tail is %v\n", *tail)
	}

	for i := 0; i < list.Length; i++ {
		if node, err := list.Get(i); err != nil {
			t.Fatalf("Get error:%v\n", err)
		} else {
			t.Logf("node is %p --> %v\n", node, *node)
		}
	}

	//if l, err := list.Get(list.Length); err != nil {
	//	t.Fatalf("Get error:%v\n", err)
	//} else {
	//	t.Logf("node is %p --> %v\n", l, *l)
	//}

	indexhead := list.GetIndex("head")
	indextail := list.GetIndex("tail")
	indexworld := list.GetIndex("world")
	t.Logf("\"head\" index is: %d\t \"tail\" index is: %d\t \"world\" index is: %d\n", indexhead, indextail, indexworld)

	indexNot := list.GetIndex("test")
	t.Logf("indexnot is: %d\n", indexNot)

	return
}

func TestDoubleListRemove(t *testing.T) {
	list := NewLinkedList()
	elements := []Element{"head", "hello", "world", "nihao", "shijie", "tail"}
	if err := list.Insert(0, elements...); err != nil {
		t.Fatalf("insert error:%v\n", err)
	}
	list.Traverse()
	i := 4
	t.Logf("############# remove %dth node ###############\n", i)
	if err := list.Remove(i); err != nil {
		t.Fatalf("Remove error:%v\n", err)
	}
	list.Traverse()

	t.Logf("############# remove head node ###############\n")
	if err := list.Pull(); err != nil {
		t.Fatalf("Pull error:%v\n", err)
	}
	list.Traverse()

	t.Logf("############# remove tail node ###############\n")
	if err := list.Pop(); err != nil {
		t.Fatalf("Pop error:%v\n", err)
	}
	list.Traverse()
}

func TestDoubleListPushAndSet(t *testing.T) {
	list := NewLinkedList()
	elements := []Element{"head", "hello", "world", "nihao", "shijie"}
	if err := list.Insert(0, elements...); err != nil {
		t.Fatalf("insert error:%v\n", err)
	}
	list.Traverse()

	t.Logf("############# Push ###############\n")
	if err := list.Push("tail"); err != nil {
		t.Fatalf("Push error:%v\n", err)
	}
	list.Traverse()

	t.Logf("############# Set ###############\n")
	if err := list.Set(1, "hi"); err != nil {
		t.Fatalf("Set error:%v\n", err)
	}
	list.Traverse()
}

func TestDoubleListMove(t *testing.T) {
	list := NewLinkedList()
	elements := []Element{"head", "hello", "world", "nihao", "shijie", "tail"}
	if err := list.Insert(0, elements...); err != nil {
		t.Fatalf("insert error:%v\n", err)
	}
	list.Traverse()

	t.Logf("############# Move ###############\n")
	nodeMove, err := list.Get(0)
	if err != nil {
		t.Fatalf("Get error:%v\n", err)
	}
	if err = list.Move(5, nodeMove); err != nil {
		t.Fatalf("Move error:%v\n", err)
	}
	list.Traverse()
}

func TestDoubleListAsList(t *testing.T) {
	list := NewLinkedList()
	elements := []Element{"head", "hello", "world", "nihao", "shijie", "tail"}
	if err := list.Insert(0, elements...); err != nil {
		t.Fatalf("insert error:%v\n", err)
	}
	list.Traverse()

	dataSet := list.AsList()
	t.Logf("data set: %v", dataSet)
}
