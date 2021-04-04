// Leetcode 300. 最长递增子序列

// Space: O(n), Time: O(n^2)
func lengthOfLIS(nums []int) int {
    n := len(nums)
    dp := make([]int, n)

    for i := 0; i < n; i++ {
        dp[i] = 1
    }

    max := 1

    for i := 1; i < n; i++ {
        for j := 0; j < i; j++ {
            if nums[i] > nums[j] && dp[i] < dp[j] + 1 {
                dp[i] = dp[j] + 1
            }
        }

        if dp[i] > max {
            max = dp[i]
        }
    }

    return max
}

// Space: O(n), Time: O(nlogn)
func lengthOfLIS(nums []int) int {
	n := len(nums)
    tails := make([]int, n)

	size := 0

	for i := 0; i < n; i++ {
		l, r := 0, size
		for l < r {
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