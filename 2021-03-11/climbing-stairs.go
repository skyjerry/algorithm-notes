// Leetcode 70. 爬楼梯

// Time: O(n), Space: O(n)
func climbStairs(n int) int {
    if n < 3 {
        return n
    }

    dp := make([]int, n + 1)
    dp[1] = 1
    dp[2] = 2

    for i := 3; i < n + 1; i++ {
        dp[i] = dp[i - 2] + dp[i - 1]
    }

    return dp[n]
}

// Time: O(n), Space: O(1)
func climbStairs(n int) int {
    if n < 3 {
        return n
    }

    one := 2
    two := 1
    ans := 0

    for i := 2; i < n; i++ {
        ans = one + two
        two = one
        one = ans
    }

    return ans
}
