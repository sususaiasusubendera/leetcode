func largestAltitude(gain []int) int {
    ans := 0
    altitude := 0
    for i := 0; i < len(gain); i++ {
        altitude += gain[i]
        ans = max(ans, altitude)
    }

    return ans
}

// array
// time: O(n)
// space: O(1)