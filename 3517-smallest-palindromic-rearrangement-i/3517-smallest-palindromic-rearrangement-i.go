func smallestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}

	temp := []byte(s)
	left, right := []byte{}, []byte{}

	if len(s)%2 == 0 {
		for i := 0; i < len(temp)/2; i++ {
			left = append(left, temp[i])
			right = append(right, temp[len(temp)-1-i])
		}

		sort.Slice(left, func(i, j int) bool {
			return left[i] < left[j]
		})

		sort.Slice(right, func(i, j int) bool {
			return right[i] > right[j]
		})

		return string(append(left, right...))
	}

	idx := 0
	for idx < len(temp)/2 {
		left = append(left, temp[idx])
		right = append(right, temp[len(temp)-1-idx])
		idx++
	}

	mid := temp[idx]

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

    sort.Slice(right, func(i, j int) bool {
		return right[i] > right[j]
	})

	ans := append(left, mid)
	ans = append(ans, right...)

    return string(ans)
}

// sorting, string
// time: O(nlog(n))
// space: O(n)