func findSmallestInteger(nums []int, value int) int {
    rFreq := map[int]int{} // map remainder -> remainder freq
    for i := 0; i < len(nums); i++ {
        // normalize a % b into (a % b + b) % b
        remainder := (nums[i] % value + value) % value
        rFreq[remainder]++
    }

    mex := 0 // mex is non-negative
    for rFreq[mex%value] > 0 {
        rFreq[mex%value]--
        mex++
    }

    return mex
}

// greedy, hash map, math
// time: O(n)
// space: O(1)