func minimumDifference(nums []int, k int) int {
    sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })

    ans := 1_000_000_007
    left, right := 0, k-1
    for right <= len(nums)-1 {
        ans = min(ans, nums[left] - nums[right])
        right++
        left++
    }

    return ans
}

// array, sliding window, sorting
// time: O(nlog(n))
// space: O(1)