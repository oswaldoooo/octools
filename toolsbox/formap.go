package toolsbox

import (
	"fmt"
	"math/rand"
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
func compareinarray[T basicmath](originarray []T, usearray map[string]struct{}, savedarray []T) []T {
	var mina, minb, minc T = 0, 0, 0
	var posa, posb int = 0, 0
	isdenfined := false
	for i := 0; i < len(originarray)/2; i++ {
		letterone := fmt.Sprintf("%v=%v", originarray[i], i)
		lettertwo := fmt.Sprintf("%v=%v", originarray[len(originarray)-i-1], len(originarray)-i-1)
		if _, ok := usearray[letterone]; !ok {
			if !isdenfined {
				mina = originarray[i]
				minb = originarray[i]
				posa = i
				posb = i
				isdenfined = true
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
				minb = originarray[len(originarray)-i-1]
				posa = len(originarray) - i - 1
				posb = len(originarray) - i - 1
				isdenfined = true
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
		if mina > minb {
			minc = mina
			mina = minb
			minb = minc
			posc := posa
			posa = posb
			posb = posc
		}
	}
	fmt.Printf("mina:%v,minb:%v\n", mina, minb)
	savedarray = append(savedarray, mina)
	savedarray = append(savedarray, minb)
	aletter := fmt.Sprintf("%v=%v", mina, posa)
	bletter := fmt.Sprintf("%v=%v", minb, posb)
	usearray[aletter] = struct{}{}
	usearray[bletter] = struct{}{}
	return savedarray
}
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
