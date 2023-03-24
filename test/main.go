package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"

	"time"

	"github.com/oswaldoooo/octools/authmethods"
	"github.com/oswaldoooo/octools/database"
	"github.com/oswaldoooo/octools/datastore"
	"github.com/oswaldoooo/octools/jwttoken"
	"github.com/oswaldoooo/octools/math"
	// "google.golang.org/appengine/runtime"
)

type user struct {
	id   string
	name string
	age  string
}

var mutx sync.Mutex

func main() {
	// var usr = user{id: "9999", name: "494724", age: "21"}
	// testreflect(usr)
	// usedb()
	testbinarysearchmap()
	// testmath()
}

func testmath() {
	testarr := []int{}
	for i := 0; i < 10; i++ {
		testarr = append(testarr, rand.Intn(1000))
	}
	// fmt.Printf("origin array>> %v\n", testarr)
	// max, min, _ := math.MaxandMin(testarr)
	ti := time.Now()
	math.MaxandMin(testarr, make(chan int), make(chan int), make(chan bool))
	usetime := time.Since(ti)
	fmt.Printf("[used time %v\n", usetime)
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

func usedb() {
	dbcontroller := database.New("user_info", "test:123456@tcp(localhost:3306)/lab")
	if dbcontroller != nil {
		// insert into user_info (id,name,password)values("oc awesome","oswaldoooo","it's great!")
		// err := dbcontroller.Insert(map[string]string{"id": "'" + template.HTMLEscapeString("oc awesome") + "'", "name": "'" + template.HTMLEscapeString("oswaldoooo") + "'", "password": "'" + template.HTMLEscapeString("it's great!") + "'"})
		err := dbcontroller.InsertEasily(map[string]string{"id": "oc awesome", "name": "oswaldoooo", "password": "it's great"})
		if err == nil {
			fmt.Println("insert data success")
			userinfo := struct{ Id, Name, Password string }{}
			//select id,name,password from user_info where id='oc awesome'
			err = dbcontroller.Get(&userinfo, "id", "'oc awesome'", "id", "name", "password")
			if err == nil {
				fmt.Println("read data>>", userinfo)
				//update user_info set name='oswaldo' where id='oc awesome'
				// err = dbcontroller.Update(map[string]string{"name": "'oswaldo'"}, "id", "'oc awesome'")
				err = dbcontroller.UpdateEasily(map[string]string{"name": "oswaldo"}, "id", "oc awesome")
				if err == nil {
					fmt.Println("update success")
					err = dbcontroller.Delete("id", "'oc awesome'")
					if err == nil {
						fmt.Println("delete data success")
					}
				}
			}
		}
		if err != nil {
			fmt.Println("error >> ", err)
		}
	}

}

// test binary search
func testbinarysearch() {
	origin_arr := []string{"oswaldo cho", "oswaldo jakson", "jakson jim", "david brown", "olina omi"}
	resarr := datastore.BinarySearch("oswaldo", origin_arr)
	if resarr == nil {
		fmt.Println("not match result")
	} else {
		fmt.Printf("search result is %v\n", resarr)
	}
}
func testbinarysearchmap() {
	origin_map := map[string]int{"oswaldo cho": 1024, "oswaldo jakson": 256, "jakson jim": 512, "david brown": 6666, "olina omi": 1000}
	resmap := datastore.BinarySearchForMap("oswaldo", origin_map)
	fmt.Printf("the result is %v \n", resmap)
}
