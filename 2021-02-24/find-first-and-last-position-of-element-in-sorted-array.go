// Leetcode 34. 在排序数组中查找元素的第一个和最后一个位置

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	first := findFirst(nums, target)
	if first == -1 {
		return []int{-1, -1}
	}
	last := findLast(nums, target)

	return []int{first, last}
}

func findFirst(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1

	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if nums[left] == target {
		return left
	}
	return -1
}

func findLast(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1

	for left < right {
		mid := left + (right-left+1)>>1
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid
		}
	}

	return left
}