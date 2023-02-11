package datastore

import "fmt"

// arraystack start
func NewStack[T basictype]() *ArrayStack[T] {
	stack := ArrayStack[T]{stack: make([]T, 5), tail: 0, head: 0, top: 1}
	return &stack
}
func (s *ArrayStack[T]) Push(value T) {
	if s.head == len(s.stack) {
		newstack := make([]T, 2*len(s.stack))
		copy(newstack, s.stack)
		s.stack = newstack
	}
	s.stack[s.top] = value
	s.head++
	s.top++
}
func (s *ArrayStack[T]) Pop() T {
	if s.head == 0 {
		return s.stack[0]
	} else {
		res := s.stack[s.head]
		s.stack[s.head] = s.stack[s.top]
		s.head--
		s.top--
		return res
	}
}
func (s *ArrayStack[T]) Peek() T {
	if s.head == 0 {
		return s.stack[0]
	} else {
		res := s.stack[s.head]
		return res
	}
}
func (s *ArrayStack[T]) PrintArray() {
	for i := 0; i < s.top; i++ {
		fmt.Printf("%v =>", s.stack[i])
	}
	fmt.Println("end")
}

// arraystack end
// arrayqueen start
func NewQueen[T basictype]() *ArrayQueen[T] {
	queen := ArrayQueen[T]{queen: make([]T, 5), head: 0, tail: -1}
	return &queen
}
func (s *ArrayQueen[T]) Add(value T) {
	if s.tail == len(s.queen) {
		newqueen := make([]T, 2*len(s.queen))
		copy(newqueen, s.queen)
		s.queen = newqueen
	}
	s.queen[s.tail+1] = value
	s.tail++
}
func (s *ArrayQueen[T]) Peek() T {
	value := s.queen[0]
	return value
}
func (s *ArrayQueen[T]) Remove() {
	langth := len(s.queen)
	newqueen := make([]T, langth)
	copy(newqueen, s.queen[1:])
	s.queen = newqueen

}
func (s *ArrayQueen[T]) PrintQueen() {
	for i := 0; i < s.tail+1; i++ {
		fmt.Printf("%v =>", s.queen[i])
	}
	fmt.Println("end")
}

//arrayqueen end
