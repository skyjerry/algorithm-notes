// Leetcode 3. 无重复字符的最长子串
// 时间复杂度和空间复杂度均为O(n)
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	lenMap := make(map[byte]int)

	// i, j are right pointer and left pointer.
	ans, i, j := 0, 0, 0

	for i < len(s) {
		if v, ok := lenMap[s[i]]; ok {
			j = Max(j, v+1)
		}

		lenMap[s[i]] = i
		ans = Max(ans, i-j+1)
		i++
	}

	return ans
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}