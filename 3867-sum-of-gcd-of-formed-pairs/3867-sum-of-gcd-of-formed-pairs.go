func gcdSum(nums []int) int64 {
    preGcd := make([]int, len(nums))
    mx := 0
    for i := 0; i < len(nums); i++ {
        mx = max(mx, nums[i])
        preGcd[i] = gcd(nums[i], mx)
    }

    sort.Ints(preGcd)

    ans := 0
    for i := 0; i < len(preGcd)/2; i++ {
        ans += gcd(preGcd[i], preGcd[len(preGcd)-1-i])
    }

    return int64(ans)
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }
    return a
}

// array, math, sorting
// time: O(nlog(n))
// space: O(n)