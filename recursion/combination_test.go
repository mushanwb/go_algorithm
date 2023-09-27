package recursion

import "testing"

func TestCombination(t *testing.T) {
	n := 4
	d := climbStairs(n)
	t.Log(d)
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	a1 := 1
	a2 := 2
	for i := 2; i < n; i++ {
		tmp := a2
		a2 = a1 + a2
		a1 = tmp
	}
	return a2
}
