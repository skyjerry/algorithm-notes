// Leetcode 264. 丑数 II

// Time: O(1), Space: O(1)
func nthUglyNumber(n int) int {
    if n <= 0 {
        return 0
    }

    i2, i3, i5 := 0, 0, 0

    nums := make([]int, 1690)

    nums[0] = 1

    for i := 1; i < n; i++ {
        nums[i] = Min(nums[i2] * 2, nums[i3] * 3, nums[i5] * 5)

        if nums[i] == nums[i2] * 2 {
            i2++
        }
        if nums[i] == nums[i3] * 3 {
            i3++
        }
        if nums[i] == nums[i5] * 5 {
            i5++
        }
    }

    return nums[n - 1]
}

func Min (a, b, c int) int {
    if a < b {
        if a < c {
            return a
        } else {
            return c
        }
    } else {
        if b < c {
            return b
        } else {
            return c
        }
    }
}
