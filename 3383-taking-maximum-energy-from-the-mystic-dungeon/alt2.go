func maximumEnergy(energy []int, k int) int {
    n := len(energy)
    ans := -1 << 31 // very small number (-1 * 2^(31))
    for i := n - k; i < n; i++ {
        sum := 0
        for j := i; j >= 0; j -= k {
            sum += energy[j]
            ans = max(ans, sum)
        }
    }

    return ans
}

// time: O(n)
// space: O(1)

func max(a, b int) int {
    if a > b { return a }

    return b
}
