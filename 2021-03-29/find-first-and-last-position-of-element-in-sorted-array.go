// Leetcode 34. 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
    n := len(nums)
    if n == 0 {
        return []int{-1, -1}
    }

    first := findFirst(nums, target)
    if first == -1 {
        return []int{-1, -1}
    }

    return []int{first, findLast(nums, target)}
}

func findFirst(nums []int, target int) int {
    l, r := 0, len(nums) - 1
    for l < r {
        mid := l + (r - l) >> 1
        if nums[mid] >= target {
            r = mid
        } else {
            l = mid + 1
        }
    }

    if nums[l] == target {
        return l
    }

    return -1
}

func findLast(nums []int, target int) int {
    l, r := 0, len(nums) - 1
    for l < r {
        mid := l + (r - l + 1) >> 1
        if nums[mid] <= target {
            l = mid
        } else {
            r = mid - 1
        }
    }

    return l
}