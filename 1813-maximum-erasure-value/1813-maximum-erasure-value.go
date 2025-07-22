func maximumUniqueSubarray(nums []int) int {
    sum, maxSum := 0, 0
    left := 0
    set := map[int]bool{}
    for right := 0; right < len(nums); right++ {
        for set[nums[right]] {
            set[nums[left]] = false
            sum -= nums[left]
            left++
        }

        set[nums[right]] = true
        sum += nums[right]
        if sum > maxSum {
            maxSum = sum
        }
    }

    return maxSum
}

// time: O(n)
// space: O(n)