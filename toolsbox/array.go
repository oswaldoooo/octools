package toolsbox

import "fmt"

type basicdata interface {
	int | string | bool | byte | float64
}

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

// transfer 1D array to map,make sure your array's data were not repeat
func ArrayToMap[T basicdata](originarr []T) (res map[T]struct{}) {
	res = make(map[T]struct{})
	for _, val := range originarr {
		res[val] = struct{}{}
	}
	return
}

func Join[T any](origin ...[]T) []T {
	if len(origin) > 0 {
		total_len := 0
		for _, ele := range origin {
			total_len += len(ele)
		}
		if total_len > 0 {
			ans := make([]T, total_len)
			start := 0
			for _, ele := range origin {
				copy(ans[start:start+len(ele)], ele)
			}
			return ans
		}
	}
	return nil
}
