package database

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/jmoiron/sqlx"
	"github.com/oswaldoooo/octools/toolsbox"
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
	args_str := strings.Join(args, ",")
	esql := fmt.Sprintf("select %v from %v", args_str, s.table_name)
	if len(pattern) > 0 && len(value) > 0 {
		esql += " where " + pattern + "=" + value
	}
	err = s.db.Get(dest, esql)
	return
}
func (s *DbController) Insert(data map[string]string) (err error) {
	argsarr := toolsbox.ExportMapKeys(data)
	valarr := []string{}
	for _, ve := range data {
		valarr = append(valarr, "'"+template.HTMLEscapeString(ve)+"'")
	}
	args_str := strings.Join(argsarr, ",")
	val_str := strings.Join(valarr, ",")
	esql := fmt.Sprintf("insert into %v (%v)values(%v)", s.table_name, args_str, val_str)
	_, err = s.db.Exec(esql)
	return
}