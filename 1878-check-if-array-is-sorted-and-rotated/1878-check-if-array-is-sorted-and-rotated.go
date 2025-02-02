func check(nums []int) bool {
    if len(nums) == 1 {
        return true
    }

    // check if nums is sorted in non-decreasing order before any rotation
    for j := 1; j < len(nums); j++ {
        if nums[j] < nums[j-1] {
            break
        }
        if j == len(nums)-1 {
            return true
        }
    }
        
    // for each rotation, check if nums is sorted in non-decreasing order
    for i := 0; i < len(nums)-1; i++ {
        nums = append(nums, nums[0])
        nums = nums[1:]
        for j := 1; j < len(nums); j++ {
            if nums[j] < nums[j-1] {
                break
            }
            if j == len(nums)-1 {
                return true
            }
        }
    }
    return false
}

// brute force
// time: O(n^2)
// space: O(n)