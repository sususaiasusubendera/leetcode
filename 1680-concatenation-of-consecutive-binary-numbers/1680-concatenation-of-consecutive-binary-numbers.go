func concatenatedBinary(n int) int {
    const MOD = 1_000_000_007

    ans := 0
    length := 0
    for i := 1; i <= n; i++ {
        // if i is the pow of 2 -> length++
        if (i & (i - 1)) == 0 {
            length++
        }

        ans = ((ans << length) + i) % MOD
    }

    return ans
}

// bit manipulation, math, simulation
// time: O(n)
// space: O(1)