package test

import "testing"

func TestSlice(t *testing.T) {
	s := []int{1, 2, 3, 4}

	for i, v := range s {
		s[i] = v + 1
	}
	t.Log(s)
}
