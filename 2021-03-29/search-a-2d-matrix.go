// Leetcode 74. 搜索二维矩阵

func searchMatrix(matrix [][]int, target int) bool {
    m, n := len(matrix), len(matrix[0])

    l, r := 0, m * n - 1

    for l < r {
        mid := l + (r - l) >> 1

        if matrix[mid/n][mid%n] < target {
            l = mid + 1
        } else if matrix[mid/n][mid%n] > target {
            r = mid
        } else {
            return true
        }
    }

    if matrix[l/n][l%n] == target {
        return true
    }

    return false
}