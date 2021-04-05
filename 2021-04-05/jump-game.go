// Leetcode 55. 跳跃游戏

// Time: O(n), Space: O(1)
func canJump(nums []int) bool {
    max := 0
    n := len(nums)

    for i := 0; i < n; i++ {
        if max < i {
            return false
        }

        if i + nums[i] > max {
            max = nums[i] + i
        }
        
        if max >= n - 1 {
            return true
        }
    }

    return false
}