func maxLength(nums []int) int {
    ans := 0
    for i := 0; i < len(nums); i++ {
        prod := 1
        g := 0
        l := 1
        for j := i; j < len(nums); j++ {
            prod *= nums[j]
            g = gcd(g, nums[j])
            l = lcm(l, nums[j])
            if prod == g * l {
                ans = max(ans, j-i+1)
            }
        }
    }
    return ans
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func lcm(a, b int) int {
    if a == 0 || b == 0 {
        return 0
    }
    return abs((a/gcd(a, b))*b)
}

// array, math
// time: O(n^2)
// space: O(1)