// Leetcode 33. Search in Rotated Sorted Array
func search(nums []int, target int) int {
	// Init array length, two ptr.
	n := len(nums)
	left, right := 0, n-1

	for left <= right {
		mid := (right-left)>>1 + left

		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] {
			if nums[left] <= target && target <= nums[mid] {
				return binarySearch(nums, left, mid, target)
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] <= target && target <= nums[right] {
				return binarySearch(nums, mid, right, target)
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

// binary Search
func binarySearch(nums []int, left, right, target int) int {
	for left <= right {
		mid := (right-left)>>1 + left
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}