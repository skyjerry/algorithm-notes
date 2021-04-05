// Leetcode 42. 接雨水

// DP. Space: O(n), Time: O(n)
func trap(height []int) int {
    n, ans := len(height), 0
    maxLeft, maxRight := make([]int, n), make([]int, n)

    for i := 1; i < n - 1; i++ {
        maxLeft[i] = Max(maxLeft[i - 1], height[i - 1])
    }

    for i := n - 2; i >= 0; i-- {
        maxRight[i] = Max(maxRight[i + 1], height[i + 1])
    }

    for i := 0; i < n - 1; i++ {
        min := Min(maxLeft[i], maxRight[i])
        if min > height[i] {
            ans += min - height[i]
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

func Min(a, b int) int {
    if a < b {
        return a
    }

    return b
}

// Two Pointers. Time: O(n), Space: O(1)
func trap(height []int) int {
    n := len(height)
    left, right := 0, n - 1
    maxLeft, maxRight := 0, 0
    ans := 0

    for left <= right {
        if maxLeft < maxRight {
            ans += Max(0, maxLeft - height[left])
            maxLeft = Max(maxLeft, height[left])
            left += 1
        } else {
            ans += Max(0, maxRight - height[right])
            maxRight = Max(maxRight, height[right])
            right -= 1
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