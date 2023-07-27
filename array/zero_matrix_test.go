package array

import (
	"testing"
)

/*
编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。
*/

/*
示例 1：
输入：
[

	[1,1,1],
	[1,0,1],
	[1,1,1]

]
输出：
[

	[1,0,1],
	[0,0,0],
	[1,0,1]

]

示例 2：
输入：
[

	[0,1,2,0],
	[3,4,5,2],
	[1,3,1,5]

]
输出：
[

	[0,0,0,0],
	[0,4,5,0],
	[0,3,1,0]

]
*/
func TestZeroMatrix(t *testing.T) {
	matrix := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	setZeroes(matrix)
	t.Log(matrix)
}

/*
m*n 的矩阵，其实就是找坐标 x,y 为0的点，找到之后，将x,y轴数据都变成0
*/
func setZeroes(matrix [][]int) {
	//定义 x,y 坐标,使用 map 存，放置数据重复
	xMap := make(map[int]int)
	yMap := make(map[int]int)
	// 找 0
	for y, v := range matrix {
		for x, m := range v {
			if m == 0 {
				// set 的实现就是基于 map,只是将 map 的value值固定
				xMap[x] = 1
				yMap[y] = 1
			}
		}
	}
	// 将对应x,y置为0
	if len(xMap) != 0 {
		for _, v := range matrix {
			for x, _ := range xMap {
				v[x] = 0
			}
		}
	}
	if len(yMap) != 0 {
		for y, _ := range yMap {
			data := matrix[y]
			for k, _ := range data {
				data[k] = 0
			}
		}
	}
}
