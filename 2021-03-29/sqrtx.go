// LeetCode 69. x 的平方根

func mySqrt(x int) int {
    if x == 0 {
        return 0
    }
    l, r, ans := 0, x, -1

    for l <= r {
        mid := l + (r - l) >> 1
        if mid * mid <= x {
            ans = mid
            l = mid + 1
        } else {
            r = mid - 1
        }
    }

    return ans
}