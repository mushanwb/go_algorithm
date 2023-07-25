package array

import "testing"

/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
请必须使用时间复杂度为 O(log n) 的算法。
*/

/*
输入: nums = [1,3,5,6], target = 5
输出: 2

输入: nums = [1,3,5,6], target = 2
输出: 1

输入: nums = [1,3,5,6], target = 7
输出: 4
*/
func TestSearchInsertIndex(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 5
	index := searchInsert(nums, target)
	t.Log(index)
}

func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	if target <= nums[start] {
		return start
	}
	if target > nums[end] {
		return end + 1
	}
	for start+1 < end {
		mid := (start + end) / 2
		midNum := nums[mid]
		if midNum > target {
			end = mid
		} else if midNum < target {
			start = mid
		} else {
			return mid
		}
	}
	return start + 1
}
