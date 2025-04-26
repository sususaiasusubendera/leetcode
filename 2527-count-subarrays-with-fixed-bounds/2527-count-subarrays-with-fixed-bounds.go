func countSubarrays(nums []int, minK int, maxK int) int64 {
    var ans int64 = 0
    var minNumIndex, maxNumIndex, badNumIndex int = -1, -1, -1
    for i, num := range nums {
        if num < minK || num > maxK {badNumIndex = i}
        if nums[i] == maxK {maxNumIndex = i}
        if nums[i] == minK {minNumIndex = i}
        ans += int64(max(0, min(minNumIndex, maxNumIndex) - badNumIndex))
    }
    return ans
}

// NOTICE ME SENPAI!!!

func min(a, b int) int {
    if a < b {return a}
    return b
}

func max(a, b int) int {
    if a > b {return a}
    return b
}