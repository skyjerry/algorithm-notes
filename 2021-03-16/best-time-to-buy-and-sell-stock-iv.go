// Leetcode 188. 买卖股票的最佳时机 IV

// Time: O(n * k), Space: O(n * k)
func maxProfit(maxK int, prices []int) int {
    n := len(prices)
    if n < 2 {
        return 0
    }

    if maxK > n / 2 {
        return maxProfitAny(prices)
    }

    var dp [][][]int

    for i := 0; i < n; i++ {
        var tmp [][]int
        for i := 0; i < maxK + 1; i++ {
            tmp = append(tmp, make([]int, 2))
        }
        dp = append(dp, tmp)
    }

    for i := 0; i < n; i++ {
        for k := maxK; k >= 1; k-- {
            if i == 0 {
                dp[i][k][0] = 0
                dp[i][k][1] = -prices[i]
            } else {
                dp[i][k][0] = Max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
                dp[i][k][1] = Max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
            }
        }
    }
    return dp[n-1][maxK][0]
}

func maxProfitAny(prices []int) int {
    n := len(prices)
    if n < 2 {
        return 0
    }

    dp_i_0 := 0
    dp_i_1 := - (1<<31 - 1)

    for i := 0; i < n; i++ {
        temp := dp_i_0
        dp_i_0 = Max(dp_i_0, dp_i_1 + prices[i])
        dp_i_1 = Max(dp_i_1, temp - prices[i])
    }

    return dp_i_0
}

func Max(a, b int) int {
    if a > b {
        return a
    }
    return b
}