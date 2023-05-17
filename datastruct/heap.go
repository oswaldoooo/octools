package datastruct

import (
	"errors"
	"fmt"
	"sync"
)

var mutx sync.Mutex

type Heap[T any] struct {
	core []T
}

func (s *Heap[T]) Size() int {
	return len(s.core)
}
func (s *Heap[T]) Top() T {
	return s.core[len(s.core)-1]
}
func (s *Heap[T]) Try_top() (ans T, err error) {
	if len(s.core) > 0 {
		ans = s.core[len(s.core)-1]
	} else {
		err = errors.New("Heap is nil")
	}
	return
}
func (s *Heap[T]) Isempty() bool {
	return len(s.core) == 0
}
func (s *Heap[T]) Pop() (ans T) {
	ans = s.core[len(s.core)-1]
	for !mutx.TryLock() {
	}
	s.core = s.core[:len(s.core)-1]
	mutx.Unlock()
	return
}
func (s *Heap[T]) Try_pop() (ans T, err error) {
	if len(s.core) > 0 {
		ans = s.core[len(s.core)-1]
		for !mutx.TryLock() {
		}
		s.core = s.core[:len(s.core)-1]
		mutx.Unlock()
	} else {
		err = errors.New("Heap is null")
	}
	return
}
func (s *Heap[T]) Push_back(obj T) {
	for !mutx.TryLock() {
	}
	s.core = append(s.core, obj)
	mutx.Unlock()
}
func (s *Heap[T]) Out() {
	for _, val := range s.core {
		fmt.Println(val)
	}
}
func NewHeap[T any]() (res *Heap[T]) {
	return new(Heap[T])
}
