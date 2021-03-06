// Leetcode 53. 最大子序和
func maxSubArray(nums []int) int {
    n := len(nums)

    // Array optimization to variable. Space O(n) => O(1).
    sum := 0
    ans := nums[0]

    for i := 0; i < n; i++ {
        if sum > 0 {
            sum += nums[i]
        } else {
            sum = nums[i]
        }

        if sum > ans {
            ans = sum
        }
    }

    return ans
}