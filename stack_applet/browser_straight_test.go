package stack_applet

import (
	"go_algorithm/tools"
	"testing"
)

/*
依次访问 A,B,C,D,E,F 五个网站
*/
func TestBrowserStraight(t *testing.T) {
	// 准备后退和前进栈
	var backS tools.StackArray
	var forthS tools.StackArray
	// 当前网站
	var now string
	// 网站集合
	dataArray := []string{"A", "B", "C", "D", "E", "F"}

	// 依次访问
	for k, v := range dataArray {
		// 往下访问的时候，才会有后退
		if k > 0 {
			// 扔后退栈
			backS.Push(now)
		}
		now = v
	}
	t.Log(backS)  //[A B C D E]
	t.Log(now)    // F
	t.Log(forthS) // [] 一直在往下访问，所以前进栈中没有元素

	// 后退
	now1 := back(&backS, &forthS, now) // 这个时候网站应该 E
	// 在后退
	now2 := back(&backS, &forthS, now1) // 这个时候网站应该 D
	// 在后退
	now3 := back(&backS, &forthS, now2) // 这个时候网站应该 C
	// 前进
	now4 := forth(&backS, &forthS, now3) // 这个时候网站应该 D
	// 后退
	now5 := back(&backS, &forthS, now4) // 这个时候网站应该 C
	// 在后退
	now6 := back(&backS, &forthS, now5) // 这个时候网站应该 B
	// 前进
	now7 := forth(&backS, &forthS, now6) // 这个时候网站应该 C
	// 在前进
	now8 := forth(&backS, &forthS, now7) // 这个时候网站应该 D
	// 在前进
	now9 := forth(&backS, &forthS, now8) // 这个时候网站应该 e
	// 在前进
	now10 := forth(&backS, &forthS, now9) // 这个时候网站应该 F

	t.Log(now1, now2, now3, now4, now5, now6, now7, now8, now9, now10)
}

// 后退
func back(backS *tools.StackArray, forthS *tools.StackArray, now string) string {
	// 后退就是先把当前页扔到前进栈
	forthS.Push(now)
	// 在从后退栈中弹出网站为当前网站
	now = backS.Pop().(string)
	return now
}

// 前进
func forth(backS *tools.StackArray, forthS *tools.StackArray, now string) string {
	// 前进就是先把当前页扔到后退栈
	backS.Push(now)
	// 在从前进栈中弹出网站为当前网站
	now = forthS.Pop().(string)
	return now
}
