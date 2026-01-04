func sumFourDivisors(nums []int) int {
    res := 0
    for _, n := range nums {
        val := sumOne(n)
        if val != -1 {
            res += val
        }
    }
    return res
}

func sumOne(n int) int {
    p := int(math.Round(math.Cbrt(float64(n))))
    if p*p*p == n && isPrime(p) {
        return 1 + p + p*p + p*p*p
    }

    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            a, b := i, n/i
            if a != b && isPrime(a) && isPrime(b) {
                return 1 + a + b + n
            }
            return -1
        }
    }
    return -1
}

func isPrime(x int) bool {
    if x < 2 {
        return false
    }
    for i := 2; i*i <= x; i++ {
        if x%i == 0 {
            return false
        }
    }
    return true
}

// NOTICE ME SENPAI!!!