func separateDigits(nums []int) []int {
	ans := []int{}
    for _, num := range nums {
        ans = iterateDigit(num, ans)
    }
    
    return ans
}

func iterateDigit(n int, a []int) []int {
	div := 1
	for n/div >= 10 {
		div *= 10
	}

	for div > 0 {
        digit := n / div
        a = append(a, digit)

        n %= div
        div /= 10
	}

    return a
}

// array
// time: O(n)
// space: O(d)