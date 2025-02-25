func numOfSubarrays(arr []int) int {
	res, modulo := 0, int(1e9+7)
	prev := [2]int{0, 0}

	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			prev = [2]int{1 + prev[0], prev[1]}
		} else {
			prev = [2]int{prev[1], 1 + prev[0]}
		}
		res = (res + prev[1]) % modulo
	}
	return res
}

// NOTICE ME SENPAI!!!