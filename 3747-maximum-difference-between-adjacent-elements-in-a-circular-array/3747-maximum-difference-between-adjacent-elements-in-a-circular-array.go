func maxAdjacentDistance(nums []int) int {
    n := len(nums)
    result := abs(nums[0] - nums[n-1])
    for i := 0; i < n-1; i++ {
        result = max(result, abs(nums[i] - nums[i+1]))
    }
    
    return result
}

// time: O(n)
// space: O(1)

func abs(n int) int {
    if n < 0 {
        return -n
    }

    return n
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}