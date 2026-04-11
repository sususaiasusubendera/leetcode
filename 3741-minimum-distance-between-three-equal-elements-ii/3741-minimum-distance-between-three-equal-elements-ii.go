func minimumDistance(nums []int) int {
    n := len(nums)
    if n < 3 {
        return -1
    }

    m := map[int][]int{}
	x := int(10e10)
    found := false
	for i := 0; i < n; i++ {
        if _, ok := m[nums[i]]; !ok {
            m[nums[i]] = []int{i}
        } else {
            m[nums[i]] = append(m[nums[i]], i)
            if nn := len(m[nums[i]]); nn >= 3 {
                a, b, c := m[nums[i]][nn-3], m[nums[i]][nn-2], m[nums[i]][nn-1]
                x = min(x, abs(a-b) + abs(b-c) + abs(a-c))
                found = true
            }
        }
    }

	if found {
        return x
    }
    return -1
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// array, hash map
// time: O(n)
// space: O(n)

// note
// copy paste "Minimum Distance Between Three Equal Elements I" solution lmao