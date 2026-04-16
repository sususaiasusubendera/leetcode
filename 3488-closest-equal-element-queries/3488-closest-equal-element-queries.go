func solveQueries(nums []int, queries []int) []int {
	n := len(nums)
	ans := make([]int, len(queries))
	for i := range ans {
		ans[i] = -1
	}

	m := map[int][]int{} // num -> arr of idx num
	for i, num := range nums {
		m[num] = append(m[num], i)
	}

	for i, q := range queries {
		if len(m[nums[q]]) > 1 {
            // bin search
			idx := -1
			currArr := m[nums[q]]
			left, right := 0, len(currArr)-1
			for left <= right {
				mid := left + ((right - left) / 2)
				if currArr[mid] == q {
					idx = mid
					break
				} else if currArr[mid] < q {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}

            // dist out of bound (i, j) -> (i - 0) + (n - j)
			if idx == 0 {
				ans[i] = min((currArr[0])+(n-currArr[len(currArr)-1]), currArr[1]-currArr[0])
			} else if idx == len(currArr)-1 {
				ans[i] = min(currArr[len(currArr)-1]-currArr[len(currArr)-2], (n-currArr[len(currArr)-1])+currArr[0])
			} else {
				ans[i] = min(currArr[idx]-currArr[idx-1], currArr[idx+1]-currArr[idx])
			}
		}
	}

    return ans
}

// array, binary search, hash map
// time: O(n + qlog(k))
// space: O(n + q)