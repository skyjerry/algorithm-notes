// Leetcode 35. 搜索插入位置

func searchInsert(nums []int, target int) int {
    n := len(nums)

    if n == 0 {
        return 0
    }

    if target > nums[n - 1] {
        return n
    }

    l, r := 0, n - 1

    for l < r {
        mid := l + (r - l) / 2

        if nums[mid] > target {
            r = mid 
        } else if nums[mid] == target {
            return mid
        } else {
            l = mid + 1
        }
    }

    return l
}