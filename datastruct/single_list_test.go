package datastruct

import "testing"

func TestInsertData(t *testing.T) {
	n := NewSingleList()
	n.InsertData(1,100)
	n.InsertData(1,10)
	n.InsertData(1,1)
	n.Traverse()
}
