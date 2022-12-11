package leetcode

// Search
// 704. Binary Search
// https://leetcode.com/problems/binary-search/
func Search(nums []int, target int) int {
	min, max := 0, len(nums)-1
	for min <= max {
		mid := (min + max) / 2
		if nums[mid] > target {
			max = mid - 1
		} else if nums[mid] < target {
			min = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
