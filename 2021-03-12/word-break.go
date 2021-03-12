// Leetcode 139. 单词拆分
// Time: O(n^2), Space: O(n)
func wordBreak(s string, wordDict []string) bool {
    n := len(s)
    m := len(wordDict)
    if n == 0 || m == 0 {
        return false
    }

    dp := make([]bool, n + 1)
    dp[0] = true
    dict := make(map[string]bool)

    for i := 0; i < m; i++ {
        dict[wordDict[i]] = true
    }

    for i := 1; i <= n; i++ {
        for j := i - 1; j > -1; j-- {
            if dict[s[j:i]] && dp[j] {
                dp[i] = true
            }
        }
    }

    return dp[n]
}