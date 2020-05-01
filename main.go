package main

import (
	ds "datastruct"
	)

func main() {
	list := ds.NewSingleList()
	list.InsertData(1,1000)
	list.InsertData(1,100)
	list.InsertData(1,10)
	list.Traverse()
}
