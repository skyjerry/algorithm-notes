// Leetcode 64. 最小路径和

// Space: O(1), Time: O(m * n)
func minPathSum(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])

    for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				if j == 0 {
					continue
				}
				grid[i][j] = grid[i][j - 1] + grid[i][j]
			} else if j == 0 {
				grid[i][j] = grid[i - 1][j] + grid[i][j]
			} else {
				grid[i][j] = min(grid[i - 1][j], grid[i][j - 1]) + grid[i][j]
			}
		}
	}

	return grid[m - 1][n - 1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}