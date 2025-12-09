func specialTriplets(nums []int) int {
    const MOD int64 = 1_000_000_007
	freqPrev, freqNext := map[int]int64{}, map[int]int64{}

    // init freqPrev
	freqPrev[nums[0]]++

    // init freqNext
	for i := 1; i < len(nums); i++ {
		freqNext[nums[i]]++
	}


	ans := int64(0)
	for j := 1; j < len(nums)-1; j++ {
        // process index j -> update freqNext: decrease nums[j] freq in freqNext
        freqNext[nums[j]]--
        if freqNext[nums[j]] == int64(0) {
            // delete if nums[j] freq in freqNext is 0
            delete(freqNext, nums[j])
        }

        target := 2 * nums[j]
		valPrev, existsPrev := freqPrev[target]
		valNext, existsNext := freqNext[target]
		if existsPrev && existsNext {
			ans = (ans + (valPrev * valNext) % MOD) % MOD
		}

        // finished processing index j -> update freqPrev: increase nums[j] freq in freqPrev
        freqPrev[nums[j]]++
	}

    return int(ans)
}

// array, hash map
// time: O(n)
// space: O(n)