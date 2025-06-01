func distributeCandies(n int, limit int) int64 {
    return starsAndTwoBars(n+2) - 3*starsAndTwoBars(n-(limit+1)+2) + 3*starsAndTwoBars(n-2*(limit+1)+2) - starsAndTwoBars(n-3*(limit+1)+2)
}

// notice me senpai?
// stars and bar
// inclusion-exclusion
// time: O(1)
// space: O(1)

// stars and bars counting with 3 children (2 bars)
func starsAndTwoBars(x int) int64 {
    if x < 0 {
        return 0
    }

    return int64((x * (x-1)) / 2)
}