func minMoves(nums []int) int {
    minEl := 9999999999
    for _, num := range nums {
        minEl = min(minEl, num)
    }

    res := 0
    for _, num := range nums {
        res += num - minEl
    }

    return res
}

// array, math
// time: O(n)
// sapce: O(1)