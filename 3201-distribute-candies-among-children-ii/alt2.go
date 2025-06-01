func distributeCandies(n int, limit int) int64 {
    result := 0
    for i := 0; i <= min(limit, n); i++ {
        if n-i > 2*limit {
            // 0 <= b <= limit, 0 <= c <= limit
            // b + c = n - a
            // b + c <= 2 * limit
            continue
        }

        // b <= limit, c <= limit, b + c = n - a = s
        // c = s - b
        // c <= limit
        // s - b <= limit
        // b >= s - limit
        // s - limit <= b <= limit
        // min(limit, s) - max(0, s - limit) + 1
        result += min(n-i, limit) - max(0, n-i-limit) + 1
    }

    return int64(result)
}
// notice me senpai?
// enumeration
// time: O(n)
// space: O(1)

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
