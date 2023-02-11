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
