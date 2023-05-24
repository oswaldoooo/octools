package database

import "fmt"

func MakeSqlUrl(address string, port int, user string, passwd string, databse string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, passwd, address, port, databse)
}
