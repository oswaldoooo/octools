package main

import (
	"fmt"

	"github.com/oswaldoooo/octools/datastore"
)

func main() {
	list := datastore.NewListQueen[string]()
	list.Add("im first in")
	list.Add("im second in")
	list.Add("im third in")
	list.PrintQueen()
	list.Remove()
	list.PrintQueen()
	list.Add("im last in")
	list.PrintQueen()
	fmt.Println(list.Peek())
}
