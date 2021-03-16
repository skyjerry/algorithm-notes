// Leetcode 122. 买卖股票的最佳时机 II

// Time: O(n), Space: O(1)
func maxProfit(prices []int) int {
    n := len(prices)
    if n < 2 {
        return 0
    }

    dp_i_1 := -prices[0]
    dp_i_0 := 0

    for i := 1; i < n; i++ {
        dp_i_0 = Max(dp_i_0, dp_i_1 + prices[i])
        dp_i_1 = Max(dp_i_1, dp_i_0 - prices[i])
    }

    return dp_i_0
}

func Max(a, b int) int {
    if a > b {
        return a
    }
    return b
}