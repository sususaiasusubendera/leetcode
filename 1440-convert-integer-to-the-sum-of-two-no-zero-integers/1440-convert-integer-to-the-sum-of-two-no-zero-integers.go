func getNoZeroIntegers(n int) []int {
    a, b := 1, n-1
    ans := []int{}
    for a < n && b > 0 {
        if !isZeroInt(a) && !isZeroInt(b) {
            ans = append(ans, a, b)
            break
        }

        a++
        b--
    }

    return ans
}

// time: O(nlog(n))
// space: O(1)

func isZeroInt(n int) bool {
    for n > 0 {
        d := n % 10
        n /= 10
        if d == 0 {
            return true
        }
    }
    return false
}