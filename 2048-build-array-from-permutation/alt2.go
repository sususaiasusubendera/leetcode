func buildArray(nums []int) []int {
    result := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        result[i] = nums[nums[i]]
    }

    return result
}

// time: O(n)
// space: O(n)
