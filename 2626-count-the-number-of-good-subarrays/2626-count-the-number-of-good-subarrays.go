func countGood(nums []int, k int) int64 {
    m := map[int]int{}
    right := 0
    countPairs := 0
    result := int64(0)
    for left := 0; left < len(nums); left++ {
        for right < len(nums) && countPairs < k {
            m[nums[right]]++
            freq := m[nums[right]]
            if freq >= 2 {
                countPairs += freq - 1
            }
            right++
        }

        if countPairs >= k {
            result += int64(len(nums) - right + 1)
        }

        m[nums[left]]--
        freq := m[nums[left]]
        if freq >= 1 {
            countPairs -= freq
        }
    }

    return result
}

// notice me senpai