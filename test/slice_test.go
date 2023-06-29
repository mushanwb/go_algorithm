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
