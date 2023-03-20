package main

import (
	"fmt"
	"reflect"

	"github.com/oswaldoooo/octools/authmethods"
	"github.com/oswaldoooo/octools/jwttoken"
)

type user struct {
	id   string
	name string
	age  string
}

func main() {
	var usr = user{id: "9999", name: "494724", age: "21"}
	testreflect(usr)
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
func testlogintemplate() {
	err := authmethods.New("user_info", "test:123456@tcp(localhost:3306)/lab")
	if err == nil {
		user := &struct {
			Id       string
			Name     string
			Password string
		}{}
		err = authmethods.CheckUser(user, "6677", "oswaldoAwesome", "id", "name", "password")
		fmt.Println("user info>>", user)
	}
	if err != nil {
		fmt.Println(err)
	}
}
func testreflect(dest interface{}) {
	tp := reflect.TypeOf(dest)
	fmt.Printf("=====%v=====", tp)
	for i := 0; i < tp.NumField(); i++ {
		fileid := tp.Field(i)

		fmt.Printf("fileid name >> %v\n", fileid.Name)
	}

}
