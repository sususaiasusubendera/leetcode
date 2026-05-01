func maxRotateFunction(nums []int) int {
    n := len(nums)
    if n < 2 {
        return 0
    }

    sum := 0
    f := 0
    for i := 0; i < n; i++ {
        sum += nums[i]
        f += i * nums[i]
    }

    ans := f
    for i := 1; i < n; i++ {
        f = (n * nums[i-1]) - sum + f
        ans = max(ans, f)
    }

    return ans
}

// array
// time: O(n)
// space: O(1)

// note (tips)
// grab some paper and a pen
// try to derive something from f(1) - f(0)