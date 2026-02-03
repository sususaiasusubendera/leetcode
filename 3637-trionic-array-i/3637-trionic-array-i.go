func isTrionic(nums []int) bool {
    p, q := false, false
    for i := 1; i < len(nums)-1; i++ {
        if !p && !q && nums[i] <= nums[i-1] {
            break
        }

        if !p && !q && nums[i-1] < nums[i] && nums[i] > nums[i+1] {
            p = true
        }

        if p && !q && nums[i-1] > nums[i] && nums[i] < nums[i+1] {
            q = true
        }

        if p && q && nums[i] >= nums[i+1] {
            p = false
            q = false
            break
        }
    }

    if p && q {
        return true
    }
    return false
}

// array
// time: O(n)
// space: O(1)