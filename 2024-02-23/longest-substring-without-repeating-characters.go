// Leetcode 3. 无重复字符的最长子串

// Space: O(1), Time: O(N)
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n < 1 {
		return n
	}

	ans, l, r := 0, 0, 0
	m := map[byte]int{}
	for r < n {
		// 如果找到之前有重复字符，则重新计算左边界是否需要右移
		// 为什么不直接使用v+1? case: abba
		if v, ok := m[s[r]]; ok {
			l = max(l, v+1)
		}
		m[s[r]] = r
		ans = max(ans, r-l+1)
		r++
	}

	return ans
}