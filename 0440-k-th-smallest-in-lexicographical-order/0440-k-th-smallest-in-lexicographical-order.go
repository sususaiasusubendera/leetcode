func findKthNumber(n int, k int) int {
    curr := 1
    k--

    for k > 0 {
        steps := countSteps(n, curr, curr+1)
        if steps <= k {
            curr += 1
            k -= steps
        } else {
            curr *= 10
            k -= 1
        }
    }

    return curr
}

// editorial hard tuning
// trie
// NOTICE ME SENPAIII!!!!!!!

func countSteps(n, prefix1, prefix2 int) int {
    steps := 0
    for prefix1 <= n {
        steps += min(n+1, prefix2) - prefix1
        prefix1 *= 10
        prefix2 *= 10
    }
    return steps
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}