func countOdds(low int, high int) int {
    count := (high - low) / 2
    if low % 2 == 1 || high % 2 == 1 { count++ }
    return count
}

// math
// time: O(1)
// space: O(1)