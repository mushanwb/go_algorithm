package tools

import "testing"

type QueueArray []interface{}

func TestQueue(t *testing.T) {
	var q QueueArray
	q.Push(1)
	q.Push(2)
	t.Log(q.Pop())
	q.Push(3)
	t.Log(q.Pop())
	t.Log(q.Pop())
	t.Log(q.Pop())
	t.Log(q.Len())
	t.Log(q)
}

func (q *QueueArray) Push(data interface{}) {
	*q = append(*q, data)
}

func (q *QueueArray) Pop() (data interface{}) {
	if q.Len() == 0 {
		return nil
	}
	data = (*q)[0]
	*q = (*q)[1:len(*q)]
	return data
}

func (q *QueueArray) Len() (n int) {
	return len(*q)
}
