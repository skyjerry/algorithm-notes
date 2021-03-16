// Leetcode 123. 买卖股票的最佳时机 III

// Time: O(n * k) = O(n), Space: O(n * k) = O(n)
func maxProfit(prices []int) int {
    n := len(prices)
    if n < 2 {
        return 0
    }

    dp := make([][][]int, n)
    maxK := 2

    for i := 0; i < n; i++ {
        tmp := make([][]int, maxK + 1)
        for j := 0; j < maxK + 1; j++ {
            tmp[j] = make([]int, 2)
        }
        dp[i] = tmp
    }

    for i := 0; i < n; i++ {
        for j := maxK; j > 0; j-- {
            if i == 0 {
                dp[i][j][0] = 0
                dp[i][j][1] = -prices[i]
            } else {
                dp[i][j][0] = Max(dp[i - 1][j][0], dp[i - 1][j][1] + prices[i])
                dp[i][j][1] = Max(dp[i - 1][j][1], dp[i - 1][j - 1][0] - prices[i])
            }
        }
    }

    return dp[n - 1][maxK][0]
}

// Time: O(n) = O(n), Space: O(1)
func maxProfit(prices []int) int {
    n := len(prices)
    if n < 2 {
        return 0
    }

    dp_i_0, dp_i_1, dp_j_0, dp_j_1 := 0, -prices[0], 0, -prices[0]

    for i := 0; i < n; i++ {
        dp_j_0 = Max(dp_j_0, dp_j_1 + prices[i])
        dp_j_1 = Max(dp_j_1, dp_i_0 - prices[i])
        dp_i_0 = Max(dp_i_0, dp_i_1 + prices[i])
        dp_i_1 = Max(dp_i_1, -prices[i])
    }

    return dp_j_0
}



func Max(a, b int) int {
    if a > b {
        return a
    }
    return b
}