func maxProduct(nums []int) int {
    sort.Ints(nums)

    return (nums[len(nums)-1] - 1) * (nums[len(nums)-2] - 1)
}

// array, sorting
// time: O(nlog(n))
// space: O(log(n))