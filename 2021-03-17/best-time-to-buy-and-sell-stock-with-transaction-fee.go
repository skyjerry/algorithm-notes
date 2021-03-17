// Leetcode 309. 最佳买卖股票时机含冷冻期

// Time: O(n), Space: O(n)
func maxProfit(prices []int) int {
    n := len(prices)
    if n < 2 {
        return 0
    }

    dp := make([][]int, n)

    for i := 0; i < n; i++ {
        dp[i] = make([]int, 3)
    }

    dp[0][0] = -prices[0]
    dp[0][1] = 0
    dp[0][2] = 0

    for i := 1; i < n; i++ {
        dp[i][0] = Max(dp[i - 1][0], dp[i - 1][1] - prices[i])
        dp[i][1] = Max(dp[i - 1][1], dp[i - 1][2])
        dp[i][2] = dp[i - 1][0] + prices[i]
    }

    return Max(dp[n - 1][1], dp[n - 1][2])
}

// Time: O(n), Space: O(1)
func maxProfit(prices []int) int {
    n := len(prices)
    if n < 2 {
        return 0
    }

    dp_i_0, dp_i_1, dp_i_2 := -prices[0], 0, 0

    for i := 1; i < n; i++ {
        tmp := dp_i_0
        dp_i_0 = Max(dp_i_0, dp_i_1 - prices[i])
        dp_i_1 = Max(dp_i_1, dp_i_2)
        dp_i_2 = tmp + prices[i]
    }

    return Max(dp_i_1, dp_i_2)
}

func Max(a, b int) int {
    if a > b {
        return a
    }

    return b
}

func Max(a, b int) int {
    if a > b {
        return a
    }

    return b
}

