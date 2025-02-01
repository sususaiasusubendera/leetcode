func isArraySpecial(nums []int) bool {
    if len(nums) == 1 {
        return true
    }

    firstEven := nums[0] % 2 == 0
    if firstEven { // if the first element is even
        for i := 1; i < len(nums); i++ {
            if i % 2 == 1 && nums[i] % 2 == 0 {
                return false
            } else if i % 2 == 0 && nums[i] % 2 == 1 {
                return false
            }
        }
    } else { // if the first element is odd
        for i := 1; i < len(nums); i++ {
            if i % 2 == 1 && nums[i] % 2 == 1 {
                return false
            } else if i % 2 == 0 && nums[i] % 2 == 0 {
                return false
            }
        }
    }

    return true
}

// time: O(n)
// space: O(1)
