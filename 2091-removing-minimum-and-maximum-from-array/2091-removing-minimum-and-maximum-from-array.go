func minimumDeletions(nums []int) int {
    if len(nums) == 1 {
        return 1
    }
    
    n := len(nums)

    minIdx, maxIdx := -1, -1
    minVal, maxVal := 100000, -100000
    for i, num := range nums {
        if num < minVal {
            minVal = num
            minIdx = i
        }

        if num > maxVal {
            maxVal = num
            maxIdx = i
        }
    }

    left := min(minIdx, maxIdx)
    right := max(minIdx, maxIdx)

    return min(right+1, n-left, left+1+n-right)
}

// array, greedy
// time: O(n)
// space: O(1)