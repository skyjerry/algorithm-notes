// 剑指 Offer 62. 圆圈中最后剩下的数字
func lastRemaining(n int, m int) int {
    ans := 0

    for i:=2; i<=n; i++ {
        ans = (ans + m) % i
    }

    return ans
}