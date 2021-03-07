// Leetcode 62. 不同路径

// Space: O(m * n), Time: O(m * n)
func uniquePaths(m int, n int) int {
    opt := make([][]int, m)
    for i:=0; i<m; i++ {
        opt[i] = make([]int, n)
    }
    
    for i:=0; i<m; i++ {
        opt[i][0] = 1
    }
    
    for j:=1; j<n; j++ {
        opt[0][j] = 1
    }
    
    for i:=1; i<m; i++ {
        for j:=1; j<n; j++ {
            opt[i][j] = opt[i-1][j] + opt[i][j-1]
        }
    }
    return opt[m-1][n-1]
}

// Space: O(2 * n) = O(n), Time: O(m * n)
func uniquePaths(m int, n int) int {
    pre := make([]int, n)
    cur := make([]int, n)
    for i:=0; i<n; i++ {
        pre[i] = 1
        cur[i] = 1
    }

    for i:=1; i<m; i++ {
        for j:=1; j<n; j++ {
            cur[j] = pre[j] + cur[j - 1]
        }
        pre = append([]int{}, cur...)
    }
    return pre[n - 1]
}

// Space: O(n), Time: O(m * n)
func uniquePaths(m int, n int) int {
    cur := make([]int, n)
    for i:=0; i<n; i++ {
        cur[i] = 1
    }

    for i:=1; i<m; i++ {
        for j:=1; j<n; j++ {
            cur[j] += cur[j - 1]
        }
    }
    return cur[n - 1]
}