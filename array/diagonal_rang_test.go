package array

import "testing"

/*
给你一个大小为 m x n 的矩阵 mat ，请以对角线遍历的顺序，用一个数组返回这个矩阵中的所有元素。

文档：https://leetcode.cn/leetbook/read/array-and-string/cuxq3/
输入：mat = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,4,7,5,3,6,8,9]

输入：mat = [[1,2],[3,4]]
输出：[1,2,3,4]
*/
func TestDiagonalRang(t *testing.T) {
	mat := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	result := findDiagonalOrder(mat)
	t.Log(result)
}

func findDiagonalOrder(mat [][]int) []int {
	var result []int
	return result
}
