func countSubarrays(nums []int, k int) int64 {
    max := -1
    for i := 0; i < len(nums); i++ {
        if nums[i] > max {
            max = nums[i]
        }
    }

    left := 0
    result := 0
    countMax := 0
    for right := 0; right < len(nums); right++ {
        if nums[right] == max {
            countMax++
        }

        for countMax >= k {
            result += len(nums) - right
            if nums[left] == max {
                countMax--
            }

            left++
        }
    }

    return int64(result)
}

// sliding window
// time: O(n)
// space: O(1)