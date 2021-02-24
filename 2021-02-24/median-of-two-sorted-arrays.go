// Leetcode 4. 寻找两个正序数组的中位数

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	left := (m + n + 1) >> 1
	right := (m + n + 2) >> 1

	return float64(findKth(nums1, nums2, left)+findKth(nums1, nums2, right)) / 2
}

func findKth(nums1, nums2 []int, k int) int {
	m := len(nums1)
	n := len(nums2)

	if m == 0 {
		return nums2[k-1]
	}
	if n == 0 {
		return nums1[k-1]
	}

	if k == 1 {
		return Min(nums1[0], nums2[0])
	}

	nums1k := Min(k/2-1, m-1)
	nums2k := Min(k/2-1, n-1)

	if nums1[nums1k] > nums2[nums2k] {
		return findKth(nums1, nums2[nums2k+1:], k-nums2k-1)
	} else {
		return findKth(nums1[nums1k+1:], nums2, k-nums1k-1)
	}

}

func Min(a, b int) int {
	if a > b {
		return b
	}

	return a
}