// LeetCode 69. x 的平方根
// https://leetcode-cn.com/problems/sqrtx/

func mySqrt(x int) int {
	left, right, ans := 0, x, -1

	for left <= right {
		mid := left + (right-left)>>1

		if mid*mid <= x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}