func smallestNumber(n int, t int) int {
    for i := n; i <= 100; i++ {
        if prodDigits(i) % t == 0 {
            return i
        }
    }

    return -1
}

func prodDigits(n int) int {
    res := 1
    for n > 0 {
        res *= n % 10
        n /= 10
    }

    return res
}

// math
// time: O(1)
// space: O(1)