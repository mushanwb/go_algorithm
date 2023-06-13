package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	a := []int{4, 1, 1, 3, 5, 2, 9, 7, 8, 10}
	bubbleSort(&a)
	fmt.Println(a)
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
