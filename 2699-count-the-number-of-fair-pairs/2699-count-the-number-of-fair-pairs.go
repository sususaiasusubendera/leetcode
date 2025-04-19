func countFairPairs(nums []int, lower int, upper int) int64 {
    ans, n, l, r := 0, len(nums), len(nums), len(nums)-1
    sort.Ints(nums)
    for i, v := range nums {
        for l-1 > i && v + nums[l-1] >= lower { l-- }
        if l <= i { l = i+1 }
        if l < n && v + nums[l] > upper { continue }
        for r > l && v + nums[r] > upper { r-- }
        if r < l || l == n { continue }
        ans += (r - l + 1)
    }
    return int64(ans)
}

// notice me senpai:(