// Leetcode 344. 反转字符串
func reverseString(s []byte)  {
    n := len(s)

    if n == 0 {
        return
    }

    for i := 0; i < n / 2; i++ {
        s[i], s[n - i - 1] = s[n - i - 1], s[i]
    }
}