func applyOperations(nums []int) []int {
    writeIndex := 0
    for i:= 0; i < len(nums); i++ {
        if i < len(nums)-1 && nums[i] == nums[i+1] && nums[i] != 0 {
            nums[i] *= 2
            nums[i+1] = 0
        }

        if nums[i] != 0 {
            if i != writeIndex {
                nums[i], nums[writeIndex] = nums[writeIndex], nums[i]
            }
            writeIndex++
        }
    }

    return nums
}

// one pass
// time: O(n)
// space: O(1)