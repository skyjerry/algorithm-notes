// Leetcode 55. 跳跃游戏
// Time: O(n), Space: O(1)
func canJump(nums []int) bool {
    max := 0
    n := len(nums)

    for i := 0; i < n; i++ {
        if max >= i && i + nums[i] > max {
            max = i + nums[i]
        }

        if max >= n - 1 {
            return true
        }
    }

    return false
}