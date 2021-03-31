// Leetcode 50. Pow(x, n)

func myPow(x float64, n int) float64 {
    if x == 0.0 {
        return 0.0
    }

    ans := 1.0

    if n < 0 {
        x = 1 / x
        n = -n
    }

    for n > 0 {
        if n & 1 == 0 {
            x = x * x
            n = n / 2
        } else {
            ans = ans * float64(x)
            n = n - 1
        }
    }

    return ans
}