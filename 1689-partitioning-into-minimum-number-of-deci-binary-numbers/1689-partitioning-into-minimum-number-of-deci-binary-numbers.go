func minPartitions(n string) int {
    m := 0
    for i := 0; i < len(n); i++ {
        m = max(m, int(n[i] - '0'))
    }

    return m
}

// greedy, string
// time: O(n)
// space: O(1)