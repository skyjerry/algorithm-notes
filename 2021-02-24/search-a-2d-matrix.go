// Leetcode 74. 搜索二维矩阵
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	left, right := 0, m*n-1

	for left < right {
		mid := left + (right-left)>>1
		if matrix[mid/n][mid%n] < target {
			left = mid + 1
		} else if matrix[mid/n][mid%n] > target {
			right = mid
		} else {
			return true
		}
	}

	if matrix[left/n][left%n] == target {
		return true
	}

	return false
}