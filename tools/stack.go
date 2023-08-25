package tools

import "testing"

type StackArray []interface{}

func TestStack(t *testing.T) {
	var s StackArray
	s.Push(1)
	s.Push(2)
	t.Log(s.Pop())
	s.Push(3)
	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Len())
	t.Log(s)
}

func (s *StackArray) Push(data interface{}) {
	*s = append(*s, data)
}

func (s *StackArray) Pop() (data interface{}) {
	if s.Len() == 0 {
		return nil
	}
	data = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return data
}

func (s *StackArray) Len() (n int) {
	return len(*s)
}
