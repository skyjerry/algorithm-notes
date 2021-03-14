// Leetcode 72. 编辑距离

// Time: O(m * n), Space: O(m * n)
func minDistance(word1 string, word2 string) int {
	n1 := len(word1)
	n2 := len(word2)

	if n1 == 0 {
		return n2
	}
	if n2 == 0 {
		return n1
	}

	dp := make([][]int, n1 + 1)

	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2 + 1)
	}

	for i := 0; i <= n1; i++ {
		dp[i][0] = i
	}
	for i := 0; i <= n2; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if word1[i - 1] == word2[j - 1] {
				dp[i][j] = dp[i - 1][j - 1]
			} else {
				dp[i][j] = min(dp[i - 1][j], dp[i][j - 1], dp[i - 1][j - 1]) + 1
			}
		}
	}

	return dp[n1][n2]
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	}
	return c
}