func countPermutations(complexity []int) int {
    n := len(complexity)
    for i := 1; i < n; i++ {
        if complexity[i] <= complexity[0] { return 0 }
    }

    // permutation: 1 * (n - 1)!
    const MOD = 1_000_000_007
    ans := 1
    for i := n - 1; i > 0; i-- {
        ans = (ans * i) % MOD
    }

    return ans
}

// math, combinatorics
// time: O(n)
// space: O(1)