// Leetcode 11. 盛最多水的容器

func maxArea(height []int) int {
    l, r, ans := 0, len(height) - 1, 0

    for l < r {
        if height[l] < height[r] {
            ans = Max(height[l] * (r - l), ans)
            l++
        } else {
            ans = Max(height[r] * (r - l), ans)
            r--
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