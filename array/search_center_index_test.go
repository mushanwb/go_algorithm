package array

import "testing"

/*
数组 中心下标 是数组的一个下标，其左侧所有元素相加的和等于右侧所有元素相加的和。

如果中心下标位于数组最左端，那么左侧数之和视为 0 ，因为在下标的左侧不存在元素。这一点对于中心下标位于数组最右端同样适用。

如果数组有多个中心下标，应该返回 最靠近左边 的那一个。如果数组不存在中心下标，返回 -1
*/

/*
输入：nums = [1, 7, 3, 6, 5, 6]
输出：3
解释：
中心下标是 3 。
左侧数之和 sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11 ，
右侧数之和 sum = nums[4] + nums[5] = 5 + 6 = 11 ，二者相等。

输入：nums = [1, 2, 3]
输出：-1
解释：
数组中不存在满足此条件的中心下标。
*/

func TestSearchCenterIndex(t *testing.T) {
	//nums := []int{1, 7, 3, 6, 5, 6}
	nums := []int{1, 2, 3}
	index := pivotIndex(nums)
	t.Log(index)
}

/*
执行用时：16 ms, 在所有 Go 提交中击败了86.74%的用户
内存消耗：6.2 MB, 在所有 Go 提交中击败了36.21% 的用户
*/
func pivotIndex(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum = sum + v
	}
	leftSum := 0
	for k, v := range nums {
		if leftSum == sum-v-leftSum {
			return k
		} else {
			leftSum = leftSum + v
		}
	}
	return -1
}
