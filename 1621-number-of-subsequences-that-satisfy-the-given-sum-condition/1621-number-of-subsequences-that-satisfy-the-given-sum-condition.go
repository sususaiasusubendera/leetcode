func numSubseq(nums []int, target int) int {
    const mod = 1e9 + 7
    sort.Ints(nums)

    // precompute power of 2
    // powers[i] = 2^i % mod
    powers := make([]int, len(nums)+1)
    powers[0] = 1
    for i := 1; i <= len(nums); i++ {
        powers[i] = (powers[i-1] * 2) % mod
    }

    result := 0
    left, right := 0, len(nums)-1
    for left <= right {
        if nums[left] + nums[right] <= target {
            result = (result + powers[right-left]) % mod
            left++
        } else {
            right--
        }
    }

    return result
}

// NOTICE ME SENPAI!!!