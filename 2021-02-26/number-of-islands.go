// Leetcode 200. 岛屿数量
func numIslands(grid [][]byte) int {
	if grid == nil {
		return 0
	}

	n := len(grid)
	m := len(grid[0])
	ans := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				ans++
			}
		}
	}

	return ans
}

func dfs(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] != '1' {
		return
	}

	grid[i][j] = '0'

	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)

}