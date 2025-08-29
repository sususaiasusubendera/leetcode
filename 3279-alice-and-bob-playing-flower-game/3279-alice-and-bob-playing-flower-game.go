func flowerGame(n int, m int) int64 {
    ans := (((n+1)/2) * (m/2)) + ((n/2) * ((m+1)/2))
    return int64(ans)
}

// math
// time: O(1)
// space: O(1)