package _map

import (
	"go_algorithm/tools"
	"testing"
)

func TestShortPath(t *testing.T) {
	// 路径
	// a -> [b, c, d]
	// b -> [c, f],  c -> [f], d -> [e, g]
	// f -> [g], e -> [f, h]
	// g -> [h]
	// 最短路径：a->d->g->h 或者 a->d->e->h
	var q tools.QueueArray
	q.Push(1)
	t.Log(q.Pop())
}
