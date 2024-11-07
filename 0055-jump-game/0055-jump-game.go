func canJump(nums []int) bool {
    maxj := 0
    for i := 0; i < len(nums); i++ {
        if i > maxj {
            return false
        }

        if i + nums[i] > maxj {
            maxj = i + nums[i]
        }

        if maxj >= len(nums)-1 {
            return true
        }
    }
    return false
}