package database

import "github.com/oswaldoooo/octools/datastruct"

type ControllerPool struct {
	core datastruct.Heap[*DbController]
}

func (s *ControllerPool) IsAvailable() bool {
	return s.core.Size() > 0
}
func (s *ControllerPool) GetConnection() *DbController {
	for !s.IsAvailable() {
	} //wait it's available
	return s.core.Pop()
}
func (s *ControllerPool) Back(origin *DbController) {
	s.core.Push_back(origin)
}
func (s *ControllerPool) Available_Number() int {
	return s.core.Size()
}
func NewPool(size int, tablename, url string) *ControllerPool {
	ans := new(ControllerPool)
	if size > 0 {
		for i := 0; i < size; i++ {
			ans.core.Push_back(New(tablename, url))
		}
	}
	return ans
}
