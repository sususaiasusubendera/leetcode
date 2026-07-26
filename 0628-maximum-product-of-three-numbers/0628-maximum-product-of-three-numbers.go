func maximumProduct(nums []int) int {
    sort.Ints(nums)

    if nums[0] < 0 && nums[1] < 0 {
        a := nums[0] * nums[1] * nums[len(nums)-1]
        b := nums[len(nums)-1] * nums[len(nums)-2] * nums[len(nums)-3]
        return max(a, b)
    } else {
        return nums[len(nums)-1] * nums[len(nums)-2] * nums[len(nums)-3]
    }
}

// sorting
// time: O(nlog(n))
// space: O(1)