package toolsbox

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type basicmath interface {
	int | uint | float32
}

func CheckArgs[T any](args []string, originmap map[string]T) bool {
	for _, v := range args {
		if _, ok := originmap[v]; !ok {
			return false
		}
	}
	return true
}
func SortArray[T basicmath](array []T) []T {
	newarray := []T{}
	usedarray := make(map[string]struct{})
	for i := 0; i < len(array)/2; i++ {
		newarray = compareinarray(array, usedarray, newarray)
		// fmt.Println(newarray)
	}
	return newarray
}

var buffstring string = ""

func compareinarray[T basicmath](originarray []T, usearray map[string]struct{}, savedarray []T) []T {
	var mina, minb T = 0, 0
	var posa, posb int = 0, 0
	isdenfined := false
	isdenfinedb := false
	for i := 0; i < len(originarray)/2; i++ {
		letterone := fmt.Sprintf("%v=%v", originarray[i], i)
		lettertwo := fmt.Sprintf("%v=%v", originarray[len(originarray)-i-1], len(originarray)-i-1)
		buffstring += "start No." + strconv.Itoa(i+1) + " compare,the mina,minb=" + fmt.Sprint(mina) + "," + fmt.Sprint(minb) + ";and this circle num is " + fmt.Sprint(originarray[i]) + "," + fmt.Sprint(originarray[len(originarray)-i-1]) + "\n"
		if _, ok := usearray[letterone]; !ok {
			if !isdenfined {
				mina = originarray[i]
				posa = i
				isdenfined = true
			} else if !isdenfinedb {
				minb = originarray[i]
				posb = i
				isdenfinedb = true
			} else {
				if originarray[i] < mina {
					minb = mina
					mina = originarray[i]
					posb = posa
					posa = i
				} else if originarray[i] < minb {
					minb = originarray[i]
					posb = i
				}
			}
		}
		if _, ok := usearray[lettertwo]; !ok {
			if !isdenfined {
				mina = originarray[len(originarray)-i-1]
				posa = len(originarray) - i - 1
				isdenfined = true
			} else if !isdenfinedb {
				minb = originarray[len(originarray)-i-1]
				posb = len(originarray) - i - 1
				isdenfinedb = true
			} else {
				if originarray[len(originarray)-i-1] < mina {
					minb = mina
					mina = originarray[len(originarray)-i-1]
					posb = posa
					posa = len(originarray) - i - 1
				} else if originarray[len(originarray)-i-1] < minb {
					minb = originarray[len(originarray)-i-1]
					posb = len(originarray) - i - 1
				}
			}
		}
	}
	buffstring += fmt.Sprintf("mina:%v,minb:%v,usedletter:%v\n", mina, minb, usearray)
	savedarray = append(savedarray, mina)
	savedarray = append(savedarray, minb)
	aletter := fmt.Sprintf("%v=%v", mina, posa)
	bletter := fmt.Sprintf("%v=%v", minb, posb)
	usearray[aletter] = struct{}{}
	usearray[bletter] = struct{}{}
	return savedarray
}

// length int,{max,min}(option)
func MakeRandArray(length int, args []int) []int {
	resarr := make([]int, length)
	if len(args) == 0 {
		for i := 0; i < length; i++ {
			rand.Seed(time.Now().UnixNano())
			resarr[i] = rand.Int()
		}
	} else if len(args) == 1 {
		for i := 0; i < length; i++ {
			rand.Seed(time.Now().UnixNano())
			resarr[i] = rand.Intn(args[0])
		}
	} else {
		for i := 0; i < length; i++ {
			rand.Seed(time.Now().UnixNano())
			resarr[i] = rand.Intn(args[0]-args[1]) + args[1]
		}
	}
	return resarr
}
