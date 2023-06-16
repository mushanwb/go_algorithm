package sort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	list := []int{5, 1, 1, 2, 0, 0}
	sortList := QuickSortV1(list)
	t.Log(sortList)
}

/*
方式一、采用数组合并的方式(会使用多余的空间，数据量多大运行会出问题)
*/
func QuickSortV1(list []int) (sortList []int) {
	if len(list) < 2 {
		return list
	}
	var left []int
	var right []int
	flag := list[0]
	for i, v := range list {
		if i == 0 {
			continue
		}
		if v > flag {
			right = append(right, v)
		} else {
			left = append(left, v)
		}
	}
	return append(append(QuickSortV1(left), flag), QuickSortV1(right)...)
}
