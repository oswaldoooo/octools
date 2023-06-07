package database

import (
	"fmt"
	"strings"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	tools "github.com/oswaldoooo/octools"
)

type DbController struct {
	db         *sqlx.DB
	table_name string
}

func New(table_name, dburl string) (controller *DbController) {
	dbcon, err := sqlx.Connect("mysql", dburl)
	if err == nil {
		controller = &DbController{db: dbcon, table_name: table_name}
	}
	return
}

// get data from mysql.example:select id,name,email from user_info where id=78;here is Get(yourstruct,"id","78","id","name","email"),if pattern is null,it'll not where
func (s *DbController) Get(dest interface{}, pattern, value string, args ...string) (err error) {
	args_str := "*"
	if len(args) > 0 {
		args_str = strings.Join(args, ",")
	}
	esql := fmt.Sprintf("select %v from %v", args_str, s.table_name)
	if len(pattern) > 0 && len(value) > 0 {
		switch strings.ToLower(value) {
		case "true", "false":
			value = strings.ToLower(value)
		default:
			value = "'" + template.HTMLEscapeString(value) + "'"
		}
		esql += " where " + pattern + "=" + value
	}
	err = s.db.Get(dest, esql)
	if tools.Mode == tools.DEBUG && tools.Logger != nil {
		tools.Logger.Println(esql)
	}
	return
}

// get data from mysql.example:select id,name,email from user_info where id=78;here is Get(yourstruct,"id","78","id","name","email"),if pattern is null,it'll not where
func (s *DbController) GetFrom(dest interface{}, table_name, pattern, value string, args ...string) (err error) {
	args_str := "*"
	if len(args) > 0 {
		args_str = strings.Join(args, ",")
	}
	esql := fmt.Sprintf("select %v from %v", args_str, table_name)
	if len(pattern) > 0 && len(value) > 0 {
		switch strings.ToLower(value) {
		case "true", "false":
			value = strings.ToLower(value)
		default:
			value = "'" + template.HTMLEscapeString(value) + "'"
		}
		esql += " where " + pattern + "=" + value
	}
	err = s.db.Get(dest, esql)
	if tools.Mode == tools.DEBUG && tools.Logger != nil {
		tools.Logger.Println(esql)
	}
	return
}

// insert into ...
func (s *DbController) InsertInto(table_name string, data map[string]string) (err error) {
	argsarr := []string{}
	valarr := []string{}
	for ke, ve := range data {
		argsarr = append(argsarr, ke)
		valarr = append(valarr, ve)
	}
	args_str := strings.Join(argsarr, ",")
	val_str := strings.Join(valarr, ",")
	esql := fmt.Sprintf("insert into %v (%v)values(%v)", table_name, args_str, val_str)
	_, err = s.db.Exec(esql)
	if tools.Mode == tools.DEBUG && tools.Logger != nil {
		tools.Logger.Println(esql)
	}
	return
}

// insert into ...
func (s *DbController) Insert(data map[string]string) (err error) {
	argsarr := []string{}
	valarr := []string{}
	for ke, ve := range data {
		argsarr = append(argsarr, ke)
		valarr = append(valarr, ve)
	}

	args_str := strings.Join(argsarr, ",")
	val_str := strings.Join(valarr, ",")
	symarr := make([]string, len(valarr))
	symarr = Fill(symarr, "?", len(valarr)) //test pass
	esql := fmt.Sprintf("insert into %v (%v)values(%v)", s.table_name, args_str, val_str)
	if tools.Mode == tools.DEBUG && tools.Logger != nil {
		tools.Logger.Println(esql)
	}
	_, err = s.db.Exec(esql, argsarr) //test pass
	return
}

// update
func (s *DbController) UpdateFrom(table_name string, setcontent map[string]string, patter, value string) (err error) {
	setarr := []string{}
	for ke, ve := range setcontent {
		setarr = append(setarr, ke+"="+ve)
	}
	esql := fmt.Sprintf("update %v set %v where %v=%v", table_name, strings.Join(setarr, ","), patter, value)
	_, err = s.db.Exec(esql)
	if tools.Mode == tools.DEBUG && tools.Logger != nil {
		tools.Logger.Println(esql)
	}
	return
}

// update
func (s *DbController) Update(setcontent map[string]string, patter, value string) (err error) {
	setarr := []string{}
	for ke, ve := range setcontent {
		setarr = append(setarr, ke+"="+ve)
	}
	esql := fmt.Sprintf("update %v set %v where %v=%v", s.table_name, strings.Join(setarr, ","), patter, value)
	_, err = s.db.Exec(esql)
	if tools.Mode == tools.DEBUG && tools.Logger != nil {
		tools.Logger.Println(esql)
	}
	return
}
func (s *DbController) DeleteFrom(table_name, patter, value string) (err error) {
	esql := fmt.Sprintf("delete from %v where %v=%v", table_name, patter, value)
	_, err = s.db.Exec(esql)
	if tools.Mode == tools.DEBUG && tools.Logger != nil {
		tools.Logger.Println(esql)
	}
	return
}
func (s *DbController) Delete(patter, value string) (err error) {
	esql := fmt.Sprintf("delete from %v where %v=%v", s.table_name, patter, value)
	_, err = s.db.Exec(esql)
	if tools.Mode == tools.DEBUG && tools.Logger != nil {
		tools.Logger.Println(esql)
	}
	return
}
func (s *DbController) Exec(query string, args ...any) (err error) {
	_, err = s.db.Exec(query, args...)
	if tools.Mode == tools.DEBUG && tools.Logger != nil {
		tools.Logger.Println(query)
	}
	return
}
