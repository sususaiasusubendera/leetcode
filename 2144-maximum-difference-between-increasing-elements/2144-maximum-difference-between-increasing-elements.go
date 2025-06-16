func maximumDifference(nums []int) int {
    result := -1
    for i := 0; i < len(nums)-1; i++ {
        for j := i+1; j < len(nums); j++ {
            if (i < j) && (nums[i] < nums[j]) {
                result = max(result, abs(nums[j]-nums[i]))
            }
        }
    }

    return result
}

// time: O(n^2)
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