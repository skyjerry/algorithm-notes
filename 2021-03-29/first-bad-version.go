// Leetcode 278. 第一个错误的版本

func firstBadVersion(n int) int {
    if n == 1 {
        return 1
    }

    l, r := 1, n

    for l < r {
        mid := l + (r - l) >> 1
        g := isBadVersion(mid)
        if g {
            r = mid
        } else {
            l = mid + 1
        }
    }

    return l
}