func maximumCount(nums []int) int {
    countNeg, countPos := 0, 0
    for i := 0; i < len(nums); i++ {
        if nums[i] < 0 {
            countNeg++
        } else if nums[i] > 0 {
            countPos++
        }
    }

    return max(countNeg, countPos)
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}