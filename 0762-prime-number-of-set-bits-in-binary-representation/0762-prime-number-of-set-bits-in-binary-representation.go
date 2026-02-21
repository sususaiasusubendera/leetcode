func countPrimeSetBits(left int, right int) int {
    ans := 0
    for n := left; n <= right; n++ {
        if isPrime(countSetBits(n)) {
            ans++
        }
    }
    return ans
}

// bit manipulation, brian kernighan’s algorithm
// time: O(nlog(n))
// space: O(1)

// brian kernighan’s algorithm
func countSetBits(n int) int {
    count := 0
    for n > 0 {
        n = n & (n - 1)
        count++
    }
    return count
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}