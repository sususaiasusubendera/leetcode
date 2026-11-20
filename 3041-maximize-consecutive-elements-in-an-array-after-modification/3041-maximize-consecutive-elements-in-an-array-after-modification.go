func maxSelectedElements(nums []int) int {
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })

    dp := map[int]int{}
    ans := 0
    for _, num := range nums {
        dp[num+1] = max(dp[num+1], dp[num]+1)
        dp[num] = max(dp[num], dp[num-1]+1)
        ans = max(ans, dp[num], dp[num+1])
    }
    
    return ans
}

// array, dp, sorting
// time: O(nlog(n))
// space: O(n)