package toolsbox

import (
	"math/rand"
	"os"
)

var ROOTPATH = os.Getenv("OCTOOLS_HOME")

// var processlog = loginit("process")
var charlist = make([]byte, 37)

func init() {
	for i := 0; i < 26; i++ {
		charlist[i] = byte('a' + i)
	}
	for i := 0; i < 10; i++ {
		charlist[26+i] = byte('0' + i)
	}
}
func RandomMake(length int) string {
	resbytes := make([]byte, length)
	for i := 0; i < length; i++ {
		resbytes = append(resbytes, charlist[rand.Intn(len(charlist))])
	}
	return string(resbytes)
}
