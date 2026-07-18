func findGCD(nums []int) int {
    s, l := 1000, 0 // smallest num, largest num
    for _, num := range nums {
        s = min(s, num)
        l = max(l, num)
    }

    return gcd(s, l)
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }
    return a
}

// array, math
// time: O(n)
// space: O(1)