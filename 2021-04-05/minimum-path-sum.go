// Leetcode 64. 最小路径和

// Space: O(1), Time: O(m * n)
func minPathSum(grid [][]int) int {
    m, n := len(grid), len(grid[0])

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i == 0 && j == 0 {
                continue
            } else if i == 0 {
                grid[0][j] += grid[0][j - 1]
            } else if j == 0 {
                grid[i][0] += grid[i - 1][0]
            } else {
                grid[i][j] += Min(grid[i - 1][j], grid[i][j - 1])
            }
        }
    }

    return grid[m - 1][n - 1]
}

func Min(a, b int) int {
    if a < b {
        return a
    }

    return b
}