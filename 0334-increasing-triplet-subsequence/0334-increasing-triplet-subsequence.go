func increasingTriplet(nums []int) bool {
    if len(nums) <= 2 {
        return false
    }   

    leftMin, rightMax := make([]int, len(nums)), make([]int, len(nums))
    
    // left preprocessing
    leftMin[0] = nums[0]
    for i := 1; i < len(nums); i++ {
        leftMin[i] = min(leftMin[i-1], nums[i])
    }

    // right preprocessing
    rightMax[len(nums)-1] = nums[len(nums)-1]
    for i := len(nums)-2; i >= 0; i-- {
        rightMax[i] = max(rightMax[i+1], nums[i])
    }

    for i := 1; i < len(nums)-1; i++ {
        if leftMin[i-1] < nums[i] && nums[i] < rightMax[i+1] {
            return true
        }
    }

    return false
}
// greedy, preprocessing
// space: O(n); time: O(n)

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}