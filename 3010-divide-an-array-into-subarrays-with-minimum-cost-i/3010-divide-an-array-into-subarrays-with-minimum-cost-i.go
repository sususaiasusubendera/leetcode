func minimumCost(nums []int) int {
    nums1 := nums[1:]
    sort.Slice(nums1, func(i, j int) bool { return nums1[i] < nums1[j] })
    return nums[0] + nums1[0] + nums1[1]
}

// array, sorting
// time: O(nlog(n))
// space: O(1)