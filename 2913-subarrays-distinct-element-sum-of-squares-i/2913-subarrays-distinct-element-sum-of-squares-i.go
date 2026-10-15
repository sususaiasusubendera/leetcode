func sumCounts(nums []int) int {
    res := 0
    for i := 0; i < len(nums); i++ {
        s := map[int]bool{}
        for j := i; j < len(nums); j++ {
            s[nums[j]] = true
            res += len(s) * len(s)
        }
    }
    return res
}

// array, hash map
// time: O(n^2)
// space: O(n)