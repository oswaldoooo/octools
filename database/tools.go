package database

import "fmt"

func MakeSqlUrl(address string, port int, user string, passwd string, databse string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, passwd, address, port, databse)
}
func Fill[T any](origin []T, val T, size int) []T {
	for i := 0; i < size; i++ {
		if i >= len(origin) {
			origin = append(origin, val)
		} else {
			origin[i] = val
		}
	}
	return origin
}
