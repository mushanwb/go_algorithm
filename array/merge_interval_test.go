package array

import (
	"sort"
	"testing"
)

/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
*/

/*

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
*/

func TestMergeInterval(t *testing.T) {
	intervals := [][]int{
		{1, 3},
		{2, 6},
		{15, 18},
		{8, 10},
	}
	result := merge(intervals)
	t.Log(result)
}

func merge(intervals [][]int) [][]int {
	var result [][]int
	// 按照二维数组第一列排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for k, v := range intervals {
		if k == 0 {
			result = append(result, v)
		} else {
			// 取出最后一个数组
			last := result[len(result)-1]
			if v[0] <= last[1] {
				// 小则合并区间
				if v[1] > last[1] {
					// 最后一个数据大才合并，否则是包含，不用变更
					result[len(result)-1][1] = v[1]
				}
			} else {
				// 大则加入
				result = append(result, v)
			}
		}
	}
	return result
}
