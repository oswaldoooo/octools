package toolsbox

import "fmt"

// need compare array,compare array
func Comparecounts[T basicmath](originlist []T, newlist []T) bool {
	var hashmap = make(map[T]int)
	var newhashmap = make(map[T]int)
	if len(originlist) != len(newlist) {
		return false
	}
	for _, v := range originlist {
		if _, ok := hashmap[v]; !ok {
			hashmap[v] = 0
		}
		hashmap[v] += 1
	}
	for _, v := range newlist {
		if _, ok := newhashmap[v]; !ok {
			newhashmap[v] = 0
		}
		newhashmap[v] += 1
	}
	if len(hashmap) != len(newhashmap) {
		fmt.Println("new array is lack args")
		return false
	}
	for k, v := range hashmap {
		if _, ok := newhashmap[k]; !ok {
			return false
		} else if newhashmap[k] != v {
			fmt.Printf("the value %v is different\n", k)
			return false
		}
	}
	return true
}
