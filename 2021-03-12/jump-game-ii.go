// Leetcode 45. 跳跃游戏 II

// Time: O(n), Space: O(1)
func jump(nums []int) int {
    n, maxPos, end, ans := len(nums), 0, 0, 0

    for i := 0; i < n - 1; i++ {
        if i + nums[i] > maxPos {
            maxPos = i + nums[i]
        }

        if i == end {
            end = maxPos
            ans++
        }
    }

    return ans
}