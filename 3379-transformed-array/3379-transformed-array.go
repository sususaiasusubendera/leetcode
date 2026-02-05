func constructTransformedArray(nums []int) []int {
    n := len(nums)
    result := make([]int, n)
    for i := 0; i < n; i++ {
        result[i] = nums[(((i+nums[i])%n)+n)%n]
    }
    return result
}

// array, simulation
// time: O(n)
// space: O(n)