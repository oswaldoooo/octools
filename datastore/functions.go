package datastore

import (
	"fmt"
)

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

// arrayqueen end
// linklist start
func NewList[T basictype](value T) *LinkList[T] {
	list := LinkList[T]{val: value}
	return &list
}
func (s *LinkList[T]) AddFirst(value T) {
	nodeadd := s
	for nodeadd.last != nil {
		nodeadd = nodeadd.last
	}
	newnode := NewList(value)
	newnode.next = nodeadd
	nodeadd.last = newnode
}
func (s *LinkList[T]) AddLast(value T) {
	nodeadd := s
	for nodeadd.next != nil {
		nodeadd = nodeadd.next
	}
	newnode := NewList(value)
	newnode.last = nodeadd
	nodeadd.next = newnode
}
func (s *LinkList[T]) AddBeforeValue(sourcevalue T, value T) {
	var nodeinfo = make(chan *LinkList[T])
	go findvaluefromlist(s, sourcevalue, nodeinfo)
	nodes := <-nodeinfo
	close(nodeinfo)
	newnode := NewList(value)
	if nodes.last != nil {
		midnode := nodes.last
		midnode.next = newnode
		newnode.next = nodes
		nodes.last = newnode
		newnode.last = midnode
	} else {
		nodes.last = newnode
		newnode.next = nodes
	}
}
func (s *LinkList[T]) AddAfterValue(sourcevalue, value T) {
	var nodeinfo = make(chan *LinkList[T])
	go findvaluefromlist(s, sourcevalue, nodeinfo)
	nodes := <-nodeinfo
	close(nodeinfo)
	newnode := NewList(value)
	if nodes.next != nil {
		midnode := nodes.next
		midnode.last = newnode
		newnode.next = midnode
		newnode.last = nodes
		nodes.next = newnode
	} else {
		nodes.next = newnode
		newnode.last = nodes
	}
}
func findvaluefromlist[T basictype](node *LinkList[T], srcval T, nodechannel chan<- *LinkList[T]) {
	if node.val == srcval {
		nodechannel <- node
	} else {
		if node.last != nil {
			go findvaluefromlist(node.last, srcval, nodechannel)
		}
		if node.next != nil {
			go findvaluefromlist(node.next, srcval, nodechannel)
		}
	}
}
func (s *LinkList[T]) DeleteFirst() {
	nodeadd := s
	if nodeadd.isAlone() || nodeadd.last == nil {
		nodeadd = nodeadd.last
		return
	} else {
		for nodeadd.last.last != nil {
			nodeadd = nodeadd.last
		}
		nodeadd.last = nodeadd.last.last
	}
}
func (s *LinkList[T]) DeleteLast() {
	nodeadd := s
	if nodeadd.isAlone() || nodeadd.next == nil {
		nodeadd = nodeadd.next
		return
	} else {
		for nodeadd.next.next != nil {
			nodeadd = nodeadd.next
		}
		nodeadd.next = nodeadd.next.next
	}
}
func (s *LinkList[T]) PrintList() {
	leftwords := ""
	rightwords := ""
	leftadd := s
	rightadd := s
	for leftadd.last != nil || rightadd.next != nil {
		if leftadd.last != nil {
			leftwords = fmt.Sprint(leftadd.last.val) + "=>" + leftwords
			leftadd = leftadd.last
		}
		if rightadd.next != nil {
			rightwords += "=>" + fmt.Sprint(rightadd.next.val)
			rightadd = rightadd.next
		}
	}
	fmt.Printf("%v %v %v\n", leftwords, s.val, rightwords)
}
func (s *LinkList[T]) isAlone() bool {
	if s.last == nil && s.next == nil {
		return true
	} else {
		return false
	}
}

// linklist end
// listqueen start
func NewListQueen[T basictype]() *ListQueen[T] {
	newqueen := &ListQueen[T]{length: 0}
	return newqueen
}
func (s *ListQueen[T]) Add(value T) {
	node := NewList(value)
	if s.queenback == nil {
		s.queenback = node
		s.queenfront = node
	} else {
		node.next = s.queenback
		s.queenback.last = node
		s.queenback = node
	}
	s.length++
}
func (s *ListQueen[T]) Peek() T {
	res := s.queenfront.val
	return res
}
func (s *ListQueen[T]) Remove() {
	if s.queenfront.isAlone() {
		s.queenfront = s.queenback.next
	} else {
		s.queenfront = s.queenfront.last
		s.queenfront.next = s.queenfront.next.next
	}
}
func (s *ListQueen[T]) PrintQueen() {
	if s.queenback != nil {
		nodeadd := s.queenback
		nodeadd.PrintList()
	}
}

//listqueen end
