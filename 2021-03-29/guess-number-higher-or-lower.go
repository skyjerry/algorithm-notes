// Leetcode 374. 猜数字大小

func guessNumber(n int) int {
    if n == 1 {
        return 1
    }
    l, r := 0, n

    for l <= r {
        mid := l + (r - l) >> 1
        g := guess(mid)
        if g == 0 {
            return mid
        } else if g < 0 {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }

    return 0
}