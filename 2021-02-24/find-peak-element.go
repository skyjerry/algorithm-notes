// Leetcode 162. 寻找峰值
func findPeakElement(nums []int) int {
	if nums == nil {
		return -1
	}

	n := len(nums)
	left, right := 0, n-1

	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}