func countBadPairs(nums []int) int64 {
    mapDiffFreq := map[int]int{} // a (nums[i] - i) to its frequency map
    totalPairs := combTwo(len(nums))
    countGoodPairs := 0
    // find total good pairs
    for i := 0; i < len(nums); i++ {
        diff := nums[i] - i
        countGoodPairs += mapDiffFreq[diff]
        mapDiffFreq[diff]++
    }

    // total bad pairs = total pairs - total good pairs
    return int64(totalPairs - countGoodPairs)
}

// time: O(n)
// space: O(n)

// UTILS
// combination n, 2
func combTwo(n int) int {
    return (n * (n - 1)) / 2
}