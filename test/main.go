package main

import (
	"fmt"

	"github.com/oswaldoooo/octools/toolsbox"
)

func main() {
	testarr := toolsbox.MakeRandArray(10, []int{10})
	fmt.Println(testarr)
	testarr = toolsbox.SortArray(testarr)
	fmt.Println(testarr)
}
