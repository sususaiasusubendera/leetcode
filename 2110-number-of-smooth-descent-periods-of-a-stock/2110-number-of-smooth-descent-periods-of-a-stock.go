func getDescentPeriods(prices []int) int64 {
    if len(prices) == 1 {
        return 1
    }

    dp := 1
    streak := 0
    for i := 1; i < len(prices); i++ {
        if prices[i-1] - prices[i] == 1 {
            streak++
        } else {
            streak = 0
        }

        dp += 1 + streak
    }

    return int64(dp)
}

// array, dp (bot-up)
// time: O(n)
// space: O(1)