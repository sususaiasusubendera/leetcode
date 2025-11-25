func smallestRepunitDivByK(k int) int {
	if k%2 == 0 || k%5 == 0 || k%10 == 0 { return -1 }

	d := countDigit(k)
	makeOne := 0 // number with only digit 1
	for i := 0; i < d; i++ {
		makeOne *= 10
		makeOne++
	}

	for {
		if makeOne%k != 0 {
			makeOne *= 10
			makeOne++
			d++
			continue
		}
		break
	}

	return d
}

func countDigit(n int) int {
	digit := 0
	for n > 0 {
		n /= 10
		digit++
	}
	return digit
}
