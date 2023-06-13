package main

import (
	"fmt"
	"testing"
)

func TestBinary(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	search := 10
	num := binarySearch(search, a, 0, len(a)-1)
	fmt.Println(num)
}

func binarySearch(search int, a []int, start int, end int) (num int) {
	if a[start] > search || a[end] < search {
		return -1
	}

	if a[start] == search {
		return start
	}

	if a[end] == search {
		return end
	}
	mid := (start + end) / 2
	if a[mid] > search {
		return binarySearch(search, a, start, mid)
	} else if a[mid] < search {
		return binarySearch(search, a, mid, end)
	} else {
		return mid
	}

}
