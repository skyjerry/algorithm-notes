// Leetcode 11. 盛最多水的容器
func maxArea(height []int) int {
    n := len(height)

    i, j, ans := 0, n - 1, 0
    
    for i < j {
        if height[i] < height[j] {
            ans = Max(ans, height[i] * (j - i))
            i++
        } else {
            ans = Max(ans, height[j] * (j - i))
            j--
        }
    }

    return ans
}

func Max(a, b int) int {
    if a > b {
        return a
    }

    return b
}