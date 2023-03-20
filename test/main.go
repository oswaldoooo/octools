package main

import (
	"fmt"

	"github.com/oswaldoooo/octools/jwttoken"
)

func main() {
	testjwttoken()
}
func testjwttoken() {
	jt := jwttoken.NewJwt()
	token, err := jt.GenerateToken(map[string]string{"usrid": "2009022126", "usrname": "oswaldo"})
	if err == nil {
		claim, err := jt.ParseToken(token)
		if err == nil {
			fmt.Println(claim.Args)
		}
	}
	if err != nil {
		fmt.Printf("error>> %v\n", err)
	} else {
		fmt.Printf("token>> %v\n", token)
	}
}
