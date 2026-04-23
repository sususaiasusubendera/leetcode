func mirrorDistance(n int) int {
    return abs(n - reverse(n))
}

func reverse(n int) int {
    res := 0
    for n > 0 {
        d := n % 10
        res = (res * 10) + d
        n /= 10
    }
    return res
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

// math
// time: log(n)
// space: O(1)