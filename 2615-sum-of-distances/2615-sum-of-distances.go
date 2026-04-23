func distance(nums []int) []int64 {
	arr := make([]int64, len(nums))
	m := map[int][]int{}
	for i, num := range nums {
		m[num] = append(m[num], i)
	}

	for _, a := range m {
		if len(a) > 1 {
			sumPos := 0
            count := 0
			for i := 1; i < len(a); i++ {
				sumPos += a[i]
                count++
			}

            sumNeg := 0
			for i := range a {
				arr[a[i]] = int64((a[i] * (i - count)) - sumNeg + sumPos)
                if i < len(a) - 1 {
                    sumNeg += a[i]
                    sumPos -= a[i+1]
                    count--
                }
			}
		}
	}

    return arr
}

// array, hash map, math, prefix sum
// time: O(n)
// space: O(n)

// tips
// go grab some paper and a pen :)