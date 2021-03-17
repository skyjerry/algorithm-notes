// Leetcode 714. 买卖股票的最佳时机含手续费

// Time: O(n), Space: O(n)
func maxProfit(prices []int, fee int) int {
    n := len(prices)
    
    if n < 2 {
        return 0
    }

    dp := make([][]int, n)

    for i := 0; i < n; i++ {
        dp[i] = make([]int, 2)
    }

    dp[0][0] = 0
    dp[0][1] = -prices[0]
    for i := 1; i < n; i++ {
        dp[i][0] = Max(dp[i - 1][0], dp[i - 1][1] + prices[i] - fee)
        dp[i][1] = Max(dp[i - 1][1], dp[i - 1][0] - prices[i])
    }

    return dp[n - 1][0]
}



// Time: O(n), Space: O(1)
func maxProfit(prices []int, fee int) int {
    n := len(prices)
    if n < 2 {
        return 0
    }

    dp_i_0 := 0
    dp_i_1 := - (1<<31 - 1)

    for i := 0; i < n; i++ {
        temp := dp_i_0
        dp_i_0 = Max(dp_i_0, dp_i_1 + prices[i])
        dp_i_1 = Max(dp_i_1, temp - prices[i] - fee)
    }

    return dp_i_0
}




func Max(a, b int) int {
    if a > b {
        return a
    }
    return b
}