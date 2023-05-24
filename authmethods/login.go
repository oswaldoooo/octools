package authmethods

import (
	"errors"
	"fmt"
	"html/template"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/oswaldoooo/octools/database"
	"github.com/oswaldoooo/octools/toolsbox"
	"golang.org/x/crypto/bcrypt"
)

var service *LogInService

// this package pack the login verfiy method,default login by usrid,password
type user struct {
	id        string
	password  string
	otherargs map[string]string
	// name string
}
type User struct {
	Id       string
	Name     string
	Password string
}
type LogInService struct {
	dburl      string
	table_name string
	dbcon      *sqlx.DB
	*database.DbController
}
type usermanager interface {
	CreateUser() error
	FindUser() error
	UpdateUser() error
}

// create user
func (s *LogInService) CreateUser(args map[string]string) (err error) {
	//check exist password args
	if passwd, ok := args["password"]; ok {
		hashres, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
		if err == nil {
			//replace the origin password
			args["password"] = string(hashres)
		}
	}
	resargs := toolsbox.ExportMapKeys(args)
	argsstr := strings.Join(resargs, ",")
	valuearr := []string{}
	for _, ve := range args {
		ve = "'" + template.HTMLEscapeString(ve) + "'"
		valuearr = append(valuearr, ve)
	}
	valuestr := strings.Join(valuearr, ",")
	esql := fmt.Sprintf("insert into %v (%v)values(%v)", s.table_name, argsstr, valuestr)
	sr, err := s.dbcon.Exec(esql)
	if err == nil {
		rows, err := sr.RowsAffected()
		if err == nil && rows <= 0 {
			err = errors.New("create user failed,unknown error")
		}
	}
	if err != nil {
		fmt.Println(esql)
	}
	return
}

// // find user
//
//	func (s *LogInService) FindUser(dest interface{}, id string, args ...string) (err error) {
//		args_str := strings.Join(args, ",")
//		esql := fmt.Sprintf("select %v from %v where id=%v", args_str, s.table_name, id)
//		// res = new(User)
//		err = s.dbcon.Get(dest, esql)
//		//debugline
//		fmt.Println("query>> ", esql)
//		return
//	}
func New(tablename, dburl string) (err error) {
	service = &LogInService{table_name: tablename, dburl: dburl}
	service.dbcon, err = sqlx.Connect("mysql", service.dburl)
	controller := database.New(service.table_name, service.dburl)
	if controller != nil {
		service.DbController = controller
	} else {
		err = errors.New("build database controller failed")
	}
	return
}

// check the user's password
func CheckUser(dest interface{}, id, password string, args ...string) (err error) {
	service.Get(dest, "id", id, args...) //get type's all method by reflect
	//debugline
	return
}
