// Leetcode 200. 岛屿数量

func numIslands(grid [][]byte) int {
    m, n := len(grid), len(grid[0])
    ans := 0

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if (grid[i][j] == '1') {
                ans++
                dfs(grid, i, j)
            }
        }
    }

    return ans
}

func dfs(grid [][]byte, i, j int) {
    m, n := len(grid), len(grid[0])

    if i < 0 || i >= m || j < 0 || j >= n {
        return
    }

    if grid[i][j] == '0' {
        return
    }

    grid[i][j] = '0'

    direction := [][]int{
        {0, -1},
        {0, 1},
        {1, 0},
        {-1, 0},
    }

    for k := 0; k < 4; k++ {
        a := direction[k][0] + i
        b := direction[k][1] + j
        if (a > -1 && b > -1 && a < m && b < n) {
            dfs(grid, a, b)
        }
    }
}