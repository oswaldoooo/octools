package datastore

type Heap[T any] struct {
	core []T
}

func (s *Heap[T]) Top() (ans T) {
	if len(s.core) > 0 {
		ans = s.core[len(s.core)-1]
	}
	return
}
func (s *Heap[T]) Pop() (ans T) {
	if len(s.core) > 0 {
		ans = s.core[len(s.core)-1]
	}
	s.core = s.core[:len(s.core)-1]
	return
}
func (s *Heap[T]) Isempty() bool {
	return len(s.core) == 0
}
func (s *Heap[T]) Push_Back(val T) {
	s.core = append(s.core, val)
}
func (s *Heap[T]) Size() int {
	return len(s.core)
}
