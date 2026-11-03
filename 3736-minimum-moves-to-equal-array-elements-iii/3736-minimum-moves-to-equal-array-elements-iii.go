func minMoves(nums []int) int {
    maxNum := 0
    for _, num := range nums {
        maxNum = max(maxNum, num)
    }
    
    res := 0
    for _, num := range nums {
        res += maxNum - num
    }

    return res
}

// array, math
// time: O(n)
// space: O(1)