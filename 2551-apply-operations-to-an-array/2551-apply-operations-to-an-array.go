func applyOperations(nums []int) []int {
    result := make([]int, len(nums))
    p := 0
    for i:= 0; i < len(nums)-1; i++ {
        if nums[i] == nums[i+1] {
            nums[i] *= 2
            nums[i+1] = 0
        }

        if nums[i] != 0 {
            result[p] = nums[i]
            p++
        }
    }

    if nums[len(nums)-1] != 0 {
        result[p] = nums[len(nums)-1]
    }

    return result
}

// time: O(n)
// space: O(n)