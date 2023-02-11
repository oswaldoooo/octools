package main

import (
	"github.com/oswaldoooo/octools/datastore"
)

func main() {
	list := datastore.NewList[string]("im first in")
	list.AddFirst("im second in")
	list.AddLast("im third in")
	list.PrintList()
	list.DeleteFirst()
	list.PrintList()
	list.AddFirst("im fourth in")
	list.PrintList()
}
