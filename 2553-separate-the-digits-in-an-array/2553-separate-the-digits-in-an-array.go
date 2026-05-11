func separateDigits(nums []int) []int {
	ans := []int{}
    for _, num := range nums {
        ans = iterateDigit(num, ans)
    }
    
    return ans
}

func iterateDigit(n int, a []int) []int {
    s := strconv.Itoa(n)
    for i := 0; i < len(s); i++ {
        a = append(a, int(s[i] - '0'))
    }

    return a
}

// array
// time: O(d)
// space: O(k)