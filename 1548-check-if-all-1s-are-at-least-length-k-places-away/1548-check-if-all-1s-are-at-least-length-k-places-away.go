func kLengthApart(nums []int, k int) bool {
    idx := 0
    for idx < len(nums) {
        if nums[idx] == 1 {
            idxStart := idx
            for idx < len(nums)-1 && nums[idx+1] != 1 {
                idx++
            }
            if idx < len(nums)-1 && nums[idx+1] == 1 {
                if idx - idxStart < k { return false }
            }
            idx++
        } else {
            idx++
        }
    }
    return true
}

// array
// time: O(n)
// space: O(1)