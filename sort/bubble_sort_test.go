package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	a := []int{4, 1, 1, 3, 5, 2, 9, 7, 8, 10}
	bubbleSort(&a)
	t.Log(a)
}

func bubbleSort(a *[]int) {
	for i := 0; i < len(*a); i++ {
		for j := 0; j < len(*a)-1-i; j++ {
			temp := (*a)[j]
			if (*a)[j] > (*a)[j+1] {
				(*a)[j] = (*a)[j+1]
				(*a)[j+1] = temp
			}
		}
	}
}

/*
传值和传引用
*/

func TestBubbleSortV2(t *testing.T) {
	a := []int{4, 1, 1, 3, 5, 2, 9, 7, 8, 10}
	fmt.Printf("原始指针的内存地址是1 %p\n", &a)
	bubbleSortV2(a)
	fmt.Printf("原始指针的内存地址是3 %p\n", &a)
	t.Log(a)
}

func bubbleSortV2(a []int) {
	fmt.Printf("原始指针的内存地址是2 %p\n", &a)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1-i; j++ {
			temp := a[j]
			if a[j] > a[j+1] {
				a[j] = a[j+1]
				a[j+1] = temp
			}
		}
	}
}
