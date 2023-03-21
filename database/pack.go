package database

import (
	"fmt"
	"strings"
	"text/template"
)

func (s *DbController) InsertEasily(data map[string]string) (err error) {
	argsarr := []string{}
	valarr := []string{}
	for ke, ve := range data {
		switch strings.ToLower(ve) {
		case "true", "false":
			ve = strings.ToLower(ve)
		default:
			ve = "'" + template.HTMLEscapeString(ve) + "'"
		}
		argsarr = append(argsarr, ke)
		valarr = append(valarr, ve)
	}
	args_str := strings.Join(argsarr, ",")
	val_str := strings.Join(valarr, ",")
	esql := fmt.Sprintf("insert into %v (%v)values(%v)", s.table_name, args_str, val_str)
	_, err = s.db.Exec(esql)
	return
}

// update
func (s *DbController) UpdateEasily(setcontent map[string]string, patter, value string) (err error) {
	setarr := []string{}
	for ke, ve := range setcontent {
		switch strings.ToLower(ve) {
		case "true", "false":
			ve = strings.ToLower(ve)
		default:
			ve = "'" + template.HTMLEscapeString(ve) + "'"
		}
		setarr = append(setarr, ke+"="+ve)
	}
	esql := fmt.Sprintf("update %v set %v where %v=%v", s.table_name, strings.Join(setarr, ","), patter, value)
	_, err = s.db.Exec(esql)
	return
}
