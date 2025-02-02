func check(nums []int) bool {
    countRotate := 0 // rotation happens when nums[i] > nums[(i+1) % len(nums)] 
    for i := 0; i < len(nums); i++ {
        if nums[i] > nums[(i+1) % len(nums)] {
            countRotate++
        }
        if countRotate > 1 { // can't be sorted in non-decreasing order using rotation
            return false
        }
    }
    return true
}

// one pass
// time: O(n)
// space: O(1)