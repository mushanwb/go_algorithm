package sort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	list := []int{5, 1, 1, 2, 0, 0}
	QuickSortV2(list, 0, len(list)-1)
	t.Log(list)
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

/*
方式二、采用数组内部调整排序
*/

func QuickSortV2(list []int, start int, end int) {
	if start >= end {
		return
	}
	//5, 9, 1, 9, 5, 3, 7, 6, 1
	flag := list[start]
	i := start
	j := end
	for start < end {
		for start < end && list[end] >= flag {
			end--
		}
		if start < end {
			list[start] = list[end]
		}
		for start < end && list[start] < flag {
			start++
		}
		if start < end {
			list[end] = list[start]
		}
	}
	list[start] = flag
	QuickSortV2(list, i, start)
	QuickSortV2(list, start+1, j)
}
