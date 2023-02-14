package toolsbox

import "fmt"

func TestSort() {
	testarr := MakeRandArray(10, []int{10})
	newarr := SortArray(testarr)
	fmt.Println("the test array ", testarr, "\nthe new array ", newarr)
	if !Comparecounts(testarr, newarr) {
		processlog.Println(buffstring)
	}
}
