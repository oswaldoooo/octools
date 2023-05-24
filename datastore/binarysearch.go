package datastore

import (
	"github.com/oswaldoooo/octools/toolsbox"
)

func BinarySearch(content string, origin_arr []string) (resarray []string) {
	resarray = []string{}
	resarray = binarysearch(content, origin_arr, 0)
	if Comparestr(content, origin_arr[0], 50) {
		resarray = append(resarray, origin_arr[0])
	}
	return
}
func binarysearch(content string, origin_arr []string, pos int) (resarray []string) {
	resarray = []string{}
	//search left child if exist
	if 2*pos+1 < len(origin_arr) {
		if Comparestr(origin_arr[2*pos+1], content, 50) {
			resarray = append(resarray, origin_arr[2*pos+1])
		}
	} else {
		return
	}
	//search right child if exist
	if 2*pos+2 < len(origin_arr) {
		if Comparestr(origin_arr[2*pos+2], content, 50) {
			resarray = append(resarray, origin_arr[2*pos+2])
		}
	} else {
		return
	}
	leftarr := binarysearch(content, origin_arr, pos*2+1)
	rightarr := binarysearch(content, origin_arr, 2*pos+2)
	if len(leftarr) > 0 {
		resarray = append(resarray, leftarr...)
	}
	if len(rightarr) > 0 {
		resarray = append(resarray, rightarr...)
	}
	return
}

// binary search for map
func BinarySearchForMap[T interface{}](content string, origin_map map[string]T) (resmap map[string]T) {
	reslist := toolsbox.ExportMapKeysAny(origin_map)
	reskeys := BinarySearch(content, reslist)
	resmap = make(map[string]T)
	for _, key := range reskeys {
		resmap[key] = origin_map[key]
	}
	return
}
