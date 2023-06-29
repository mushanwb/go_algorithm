package test

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	s := []int{1, 2, 3, 4}

	for i, v := range s {
		s[i] = v + 1
	}
	t.Log(s)
}

// https://coolshell.cn/articles/21128.html 切片入坑文档
/*
 切片是共享内存的
*/

// 测试切片切割
func TestSliceV2(t *testing.T) {
	s := []int{1, 2, 3, 4}
	fmt.Printf("Address of i=%d:\t%p\n", s, &s[0])
	s[0] = 100
	fmt.Printf("Address of i=%d:\t%p\n", s, &s[0])

	s = s[1:4]
	fmt.Printf("Address of i=%d:\t%p\n", s, &s[0])
	s[0] = 3
	s = append(s, 5)
	t.Log(s)
}

// 测试切片后，在 append
func TestSliceV3(t *testing.T) {
	s := []int{1, 2, 3, 4}
	// 切割 s
	b := s[0:2]
	//更改 b 第一个元素,s 也会跟着改
	b[0] = 2
	// 打印 s
	t.Log(s) // [2 2 3 4]
	t.Log(b) // [2 2]
	//s 增加一个元素
	s = append(s, 5)
	//在更改b中元素
	b[0] = 3
	// 打印 s
	t.Log(s) // [2 2 3 4 5]
	t.Log(b) // [3 2]
	// s 没有变，因为 append 会复制另外一个内存，导致 s b 不在共享内存
}
