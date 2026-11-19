func minimizeSum(nums []int) int {
    n := len(nums)

    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })

    ans := min(nums[n-1] - nums[2], nums[n-3] - nums[0])
    return min(ans, nums[n-2] - nums[1])
}

// array, greedy, sorting
// time: O(nlog(n))
// space: O(1)