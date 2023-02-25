package toolsbox

import (
	"strconv"
	"strings"
)

// get number of string. example:199 $/mon,we will get 199
func getnum(origin string) (string, bool) {
	var err error
	res := []string{}
	startrecord := false
	existspot := false
	record := false
	for i := 0; i < len(origin); i++ {
		ie := origin[i : i+1]
		_, err = strconv.Atoi(ie)
		if err == nil {
			res = append(res, ie)
			startrecord = true
			if record {
				//example:199hello89,it's not number format
				return "", false
			}
		} else {
			if startrecord && ie == "." && !existspot {
				res = append(res, ie)
				existspot = true
			} else if startrecord && ie == "." && existspot {
				//example 199.99.67,it's not number format
				return "", false
			} else if startrecord {
				//example:1999str
				if !record {
					//tell computer start record number
					record = true
				}
			}
		}
	}
	finalres := strings.Join(res, "")
	return finalres, true
}

// return int,float,isfloat,isexsit
func FindNumberFromString(origin string) (int, float64, bool, bool) {
	res, ok := getnum(origin)
	if ok {
		if strings.ContainsRune(res, '.') {
			resnum, err := strconv.ParseFloat(res, 10)
			if err != nil {
				return 0, 0, false, false
			}
			return 0, resnum, true, true
		} else {
			resnum, err := strconv.Atoi(res)
			if err != nil {
				return 0, 0, false, false
			}
			return resnum, 0, false, true
		}
	} else {
		return 0, 0, false, false
	}
}
