func rotatedDigits(n int) int {
	ans := 0
	for i := 1; i <= n; i++ {
        if good(i) {
            ans++
        }
	}

    return ans
}

func good(x int) bool {
    valid := false
    for x > 0 {
        d := x % 10

        if d == 3 || d == 4 || d == 7 {
            return false
        }

        if d == 2 || d == 5 || d == 6 || d == 9 {
            valid = true
        }

        x /= 10
    }

    return valid
}

// brute force
// time: O(nlog(n))
// space: O(1)