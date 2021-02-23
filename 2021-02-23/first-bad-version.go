// Leetcode 278. 第一个错误的版本
// https://leetcode-cn.com/problems/first-bad-version/
func firstBadVersion(n int) int {
	left, right := 1, n
	for left < right {
		mid := left + (right-left)>>1
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}