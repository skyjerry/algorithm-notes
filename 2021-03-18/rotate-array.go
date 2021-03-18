// Leetcode 189. 旋转数组

// Time: O(n), Space: O(n)
func rotate(nums []int, k int)  {
    if k == 0 {
        return
    }
    n := len(nums)

    newNums := make([]int, n)

    for i := 0; i < n; i++ {
        newIndex := (i + k) % n
        newNums[newIndex] = nums[i]
    }

    for i := 0; i < n; i++ {
        nums[i] = newNums[i]
    }
}

// Time: O(n), Space: O(1)
func rotate(nums []int, k int)  {
    if k == 0 {
        return
    }
    n := len(nums)
    k = k % n

    // reverse all.
    reverse(nums, 0, n)
    reverse(nums, 0, k)
    reverse(nums, k, n)
}

func reverse(nums []int, start, end int) {
    mid := start + (end - start) >> 1
    n := start + end
    for i := start; i < mid; i++ {
        nums[i], nums[n - i - 1] = nums[n - i - 1], nums[i]
    }
}