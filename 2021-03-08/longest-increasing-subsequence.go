// Leetcode 300. 最长递增子序列

// Space: O(n), Time: O(nlogn)
func lengthOfLIS(nums []int) int {
	n := len(nums)
    tails := make([]int, n)

	size := 0

	for i := 0; i < n; i++ {
		l, r := 0, size
		for l != r {
			mid := l + (r - l) >> 1
			if tails[mid] < nums[i] {
				l = mid + 1
			} else {
				r = mid
			}
		}

		tails[l] = nums[i]
		if l == size {
			size++
		}
	}

	return size
}