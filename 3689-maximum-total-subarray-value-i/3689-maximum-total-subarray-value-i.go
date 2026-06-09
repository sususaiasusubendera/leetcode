func maxTotalValue(nums []int, k int) int64 {
    if len(nums) <= 1 {
        return int64(0)
    }

    minNum := int(^uint(0) >> 1) // max int
    maxNum := -int(^uint(0) >> 1) - 1 // min int
    for i := 0; i < len(nums); i++ {
        minNum = min(minNum, nums[i])
        maxNum = max(maxNum, nums[i])
    }

    return int64(k * (maxNum - minNum))
}

// array, greedy
// time: O(n)
// space: O(1)