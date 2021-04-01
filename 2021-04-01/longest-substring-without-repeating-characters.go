// Leetcode 3. 无重复字符的最长子串

func lengthOfLongestSubstring(s string) int {
    n := len(s)

    if n == 0 {
        return 0
    }

    i, j, ans := 0, 0, 0
    hash := make(map[byte]int)

    for j < n {
        if val, ok := hash[s[j]]; ok {
            i = Max(i, val + 1)
        }

        hash[s[j]] = j
        ans = Max(ans, j - i + 1)
        j++
    }

    return ans
}

func Max(a, b int) int {
    if a > b {
        return a
    }

    return b
}