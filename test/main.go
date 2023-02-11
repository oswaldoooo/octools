package main

import (
	"fmt"

	"github.com/oswaldoooo/octools/datastore"
)

func main() {
	queen := datastore.NewQueen[string]()
	queen.Add("im first in")
	queen.Add("im second in")
	queen.Add("im third in")
	fmt.Println(queen.Peek())
	queen.Remove()
	fmt.Println(queen.Peek())
}
