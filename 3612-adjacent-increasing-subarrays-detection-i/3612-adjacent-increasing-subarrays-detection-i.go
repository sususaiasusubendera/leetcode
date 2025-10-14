func hasIncreasingSubarrays(nums []int, k int) bool {
	leftRight := map[int]int{} // idx left -> idx right
	keys := []int{}            // stores idx left
	left := 0
	for right := 1; right < len(nums); right++ {
		if nums[right] > nums[right-1] {
			continue
		} else {
			// only stores increasing subarrays with length >= k
			if right-left >= k {
				keys = append(keys, left)
				leftRight[left] = right - 1
			}
			left = right
		}
	}
	// stores the last increasing subarray (with length >= k)
	if len(nums)-left >= k {
		keys = append(keys, left)
		leftRight[left] = len(nums) - 1
	}

	// if there is only 1 subarray: check if the subarray has length >= 2 * k
	if len(keys) == 1 && leftRight[keys[0]]-keys[0]+1 >= 2*k {
		return true
	}

	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] }) // sort, ascending

	for i := 1; i < len(keys); i++ {
		leftA, rightA := keys[i-1], leftRight[keys[i-1]]
		leftB, rightB := keys[i], leftRight[keys[i]]
		// check if subarray a & b are adjacent (all the subarrays are guaranteed to have length >= k)
		if rightA == leftB-1 {
			return true
		}

		// check if 1 subarray has length >= 2 * k
		if rightA-leftA+1 >= 2*k { return true }
		if i == len(keys)-1 {
			if rightB-leftB+1 >= 2*k { return true }
		}
	}

	return false
}

// array, hash map, sorting, two pointers
// time: O(n log(n))
// space: O(n)