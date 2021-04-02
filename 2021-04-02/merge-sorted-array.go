// Leetcode 88. 合并两个有序数组

// Space: O(1), Time: O(n)
func merge(nums1 []int, m int, nums2 []int, n int)  {
    if n == 0 {
        return
    }

    i, j, pos := m - 1, n - 1, m + n - 1

    for i > -1 && j > -1 {
        if nums2[j] > nums1[i] {
            nums1[pos] = nums2[j]
            j--
        } else {
            nums1[pos] = nums1[i]
            i--
        }
        pos--
    }

    for j > -1 {
        nums1[pos] = nums2[j]
        pos--
        j--
    }
}