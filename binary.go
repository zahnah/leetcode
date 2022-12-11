package leetcode

// FirstBadVersion
// 278. First Bad Version
// https://leetcode.com/problems/first-bad-version/
func FirstBadVersion(n int) int {
	min, max := 1, n
	for min < max {
		mid := (min + max) / 2
		if isBadVersion(mid) {
			max = mid
		} else {
			min = mid + 1
		}
	}
	return min
}

func isBadVersion(n int) bool {
	if n >= 4 {
		return true
	} else {
		return false
	}
}

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
