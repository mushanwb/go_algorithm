package _map

import (
	"go_algorithm/tools"
	"testing"
)

/*
图算法，找到路径的数量，其实就是在找起始点和重点的最短层数
*/
func TestShortPath(t *testing.T) {
	// 路径
	// a -> [b, c, d]
	// b -> [c, f],  c -> [f], d -> [e, g]
	// f -> [g], e -> [f, h]
	// g -> [h]
	// 最短路径：a->d->g->h 或者 a->d->e->h

	// 开始和结束
	start := "a"
	end := "h"

	// 将每个元素相邻的元素都放入集合
	m := make(map[string][]string)
	m["a"] = []string{"b", "c", "d"}
	m["b"] = []string{"c", "f"}
	m["c"] = []string{"f"}
	m["d"] = []string{"e", "g"}
	m["f"] = []string{"g"}
	m["e"] = []string{"f", "h"}
	m["g"] = []string{"h"}
	m["h"] = []string{}

	// 记录哪些元素已经计算过了
	l := make(map[string]int)

	// 层数
	i := 1
	// 下一层的元素个数
	fnum := 0
	// 当前层的元素个数
	nnum := 0
	// 当前层第几个元素
	dnum := 0
	// 元素队列
	var q tools.QueueArray

	// 将第一层相邻的元素都放入队列中，并计算第一层的元素个数
	for _, k := range m[start] {
		q.Push(k)
		l[k] = i
		nnum++
	}
	// 循环队列
	for q.Len() > 0 {
		f := q.Pop().(string)
		// 当前层的的元素已经遍历完，开始初始化下一层
		if dnum == nnum {
			i++
			dnum = 0
			nnum = fnum
			fnum = 0
		}
		// 如果到达重点，就退出
		if f == end {
			break
		}
		dnum++
		if v, ok := m[f]; ok {
			for _, k := range v {
				if _, okk := l[k]; !okk {
					q.Push(k)
					fnum++
					l[k] = i + 1
				}
			}
		} else {
			continue
		}
	}
	t.Log(i)
}
