package array

import "testing"

/*
给你一幅由 N × N 矩阵表示的图像，其中每个像素的大小为 4 字节。请你设计一种算法，将图像旋转 90 度。
不占用额外内存空间能否做到？
*/

/*
示例 1:
给定 matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

原地旋转输入矩阵，使其变为:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]

示例 2:
给定 matrix =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
],

原地旋转输入矩阵，使其变为:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]

*/

func TestRotationMatrix(t *testing.T) {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(matrix)
	t.Log(matrix)
}

/*
采用旋转公式，坐标 (x,y) 饶 (0,0) 旋转90度后坐标为 (y,-x),由于是原地旋转，所以需要上移，坐标为(y,-x+n-1)
按照这个公式，
方案一、可以循环遍历里面的数据，然后一个个拿出来放到另外一个数组对应的位置即可，但是会消耗额外空间
方案二、从第一个数据拿起，放到对应位置后，在把对应位置的数据拿起，放置对应位置，直到所有数据都放完
采用方案二试下
*/
func rotate(matrix [][]int) {
	// 从最外层开始遍历，先把最外层的都旋转，然后内层一次旋转
	nowLevel := 1
	// 总数
	n := len(matrix)
	// 数组层数计算（个数）/ 2，为层数，为基数需要 +1，但是最后一层不需要旋转，所以最终不用 + 1
	totalLevel := n / 2
	//最后一层结束
	for nowLevel <= totalLevel {
		// 当前层开始的索引
		nowLevelStartIndex := nowLevel - 1
		// 当前层结束的索引
		nowLevelEndIndex := n - nowLevel
		// 循环当前层要移动的数据
		for i := nowLevelStartIndex; i < nowLevelEndIndex; i++ {
			// 通过坐标拿旋转的数据,当前层是 x,当前层索引是 y（和上面公式是反的，但是不影响）
			// 拿出当前旋转的点
			x := nowLevelStartIndex
			y := i
			var data int
			data = matrix[x][y]
			// 因为是正方形，所以一个点只会旋转 4 次
			for m := 0; m <= 3; m++ {
				// 需要旋转的位置
				tmpX := x
				x = y
				y = -tmpX + n - 1
				tmpData := matrix[x][y]
				matrix[x][y] = data
				data = tmpData
			}
		}
		nowLevel++
	}
}
