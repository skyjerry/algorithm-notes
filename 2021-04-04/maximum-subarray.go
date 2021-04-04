// Leetcode 53. 最大子序和

// Space: O(n), Time: O(n)
func maxSubArray(nums []int) int {
    n := len(nums)
    dp := make([]int, n)
    dp[0] = nums[0]
    max := dp[0]

    for i := 1; i < n; i++ {
        dp[i] = Max(dp[i - 1] + nums[i], nums[i])
        max = Max(dp[i], max)
    }

    return max
}

func Max(a, b int) int {
    if a > b {
        return a
    }

    return b
}



func maxSubArray(nums []int) int {
    n := len(nums)

    // Array optimization to variable. Space O(n) => O(1).
    sum := 0
    ans := nums[0]

    for i := 0; i < n; i++ {
        if sum > 0 {
            sum += nums[i]
        } else {
            sum = nums[i]
        }

        if sum > ans {
            ans = sum
        }
    }

    return ans
}