func maxSubarraySum(nums []int, k int) int64 {
    const inf = int64(math.MaxInt64)

    best := make([]int64, k) // stores best sum for each mod
    for i := 0; i < k; i++ {
        best[i] = inf // initiate with the largest int64 number
    }
    best[0] = 0 // prefix sum = 0 at index 0

    ans := -inf
    prefix := int64(0)
    for i := 1; i <= len(nums); i++ {
        prefix += int64(nums[i-1])
        mod := i % k // i: prefix sum's index

        if best[mod] != inf {
            // sum[l..r] = P[r+1] - P[l]
            // P[r+1]: prefix
            // P[l]: best[mod]
            if prefix - best[mod] > ans {
                ans = prefix - best[mod]
            }
        }

        if prefix < best[mod] {
            best[mod] = prefix
        }
    }

    return ans
}

// array, prefix sum
// time: O(n)
// space: O(k)