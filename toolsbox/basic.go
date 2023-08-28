package toolsbox

import (
	"encoding/base32"
	"math/rand"
	"os"
	"strconv"
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
func RandomString(n int) string {
	ans := make([]byte, n)
	start := 0
	for n >= 32 {
		base32.StdEncoding.Encode(ans[start:start+32], []byte(strconv.Itoa(rand.Int())))
		start += 32
		n -= 32
	}
	if n > 0 {
		copy(ans[start:], []byte(base32.StdEncoding.EncodeToString([]byte(strconv.Itoa(rand.Int()))))[:len(ans)-start])
	}
	return string(ans)
}
