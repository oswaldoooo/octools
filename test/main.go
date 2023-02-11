package main

import (
	"github.com/oswaldoooo/octools/datastore"
)

func main() {
	list := datastore.NewList("im first in")
	list.AddFirst("im second in")
	list.AddLast("im third in")
	list.PrintList()
	list.DeleteFirst()
	list.PrintList()
	list.AddFirst("im fourth in")
	list.PrintList()
	list.DeleteLast()
	list.AddLast("im last in")
	list.PrintList()
}
