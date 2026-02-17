func reverseBits(n int) int {
	bits := make([]int, 32)
	i := 0
	for n > 0 {
		bits[i] = n % 2
		n /= 2
		i++
	}

	ans := 0
	mul := 1
	for i := 0; i < 31; i++ {
		mul *= 2
	}
	for i := 0; i < len(bits); i++ {
		ans += bits[i] * mul
		mul /= 2
	}

	return ans
}

// bit manipulation
// time: O(1)
// space: O(1)