// Leetcode 215. 数组中的第K个最大元素

// Space: O(1), Time: O(N)
func findKthLargest(nums []int, k int) int {
    n := len(nums)

    target := n - k
    left, right := 0, n - 1
    
    for {
        index := partition(nums, left, right)

        if index == target {
            return nums[index]
        } else if index < target {
            left = index + 1
        } else {
            right = index - 1
        }
    }

    return -1
}

func partition(nums []int, left, right int) int {
    i, j := left, left

    for j < right {
        if nums[j] < nums[right] {
            nums[i], nums[j] = nums[j], nums[i]
            i++
        }
        j++
    }

    nums[i], nums[right] = nums[right], nums[i]

    return i
}