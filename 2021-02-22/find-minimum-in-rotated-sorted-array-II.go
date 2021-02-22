// Leetcode 154. Find Minimum in Rotated Sorted Array II

func findMin(nums []int) int {
	n := len(nums)

	// Init two ptr.
	left, right := 0, n-1

	for left < right {
		mid := (right-left)>>1 + left
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			// if mid == right, shrink right
			right--
		}
	}

	return nums[left]
}