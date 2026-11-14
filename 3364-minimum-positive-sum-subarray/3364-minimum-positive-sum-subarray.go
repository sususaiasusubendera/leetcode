func minimumSumSubarray(nums []int, l int, r int) int {
    n := len(nums)

    pre := make([]int, n+1)
    for i := 0; i < n; i++ {
        pre[i+1] = pre[i] + nums[i]
    }

    ans := 9999999999
    for size := l; size <= r; size++ {
        for i := size; i <= n; i++ {
            sum := pre[i]-pre[i-size]
            if sum > 0 && sum < ans {
                ans = sum
            }
        }
    }

    if ans == 9999999999 {
        return -1
    }

    return ans
}

// array, prefix sum
// time: O(n^2)
// space: O(n)