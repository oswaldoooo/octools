package main

import (
	"fmt"

	"github.com/oswaldoooo/octools/datastore"
)

func main() {
	stack := datastore.NewListStack[string]()
	stack.Push("im first in")
	stack.Push("im second in")
	stack.Push("im third in")
	stack.PrintStack()
	fmt.Println(stack.Peek().GetVal())
	fmt.Println(stack.Pop().GetVal())
	stack.PrintStack()
	stack.Push("im last in")
	stack.PrintStack()
}
