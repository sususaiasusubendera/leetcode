func isGood(nums []int) bool {
    if len(nums) <= 1 {
        return false
    }

    sort.Ints(nums)

    if nums[len(nums)-2] != nums[len(nums)-1] || nums[len(nums)-1] != len(nums)-1 {
        return false
    }

    for i, num := range nums {
        if i != len(nums) - 1 {
            if num != i + 1 {
                return false
            }
        }
    }

    return true
}

// array, sorting
// time: O(nlog(n))
// space: O(1)