package datastore

type basictype interface {
	int | string | bool | byte | uint | int64 | int32 | float32 | float64
}
type ArrayStack[T basictype] struct {
	stack []T
	head  int
	tail  int
	top   int
}
type ArrayQueen[T basictype] struct {
	queen      []T
	head, tail int
}
type LinkList[T basictype] struct {
	last *LinkList[T]
	next *LinkList[T]
	val  T
}
type ListQueen[T basictype] struct {
	queenback  *LinkList[T]
	queenfront *LinkList[T]
	length     int
}
