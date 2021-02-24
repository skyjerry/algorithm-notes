// Leetcode 35. 搜索插入位置
func searchInsert(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	left, right := 0, n-1

	for left+1 < right {
		mid := left + (right-left)>>1

		if nums[mid] > target {
			right = mid
		} else {
			left = mid
		}
	}

	if nums[left] >= target {
		return left
	} else if nums[right] >= target {
		return right
	} else if nums[right] < target {
		return right + 1
	}

	return -1
}