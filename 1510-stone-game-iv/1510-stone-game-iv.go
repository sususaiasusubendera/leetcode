func winnerSquareGame(n int) bool {
    memo := make([]int, n+1) // -1: not visited, 0: lose, 1: win
    for i := range memo {
        memo[i] = -1
    }
    
    var dp func(n int) bool
    dp = func(n int) bool {
        if n == 0 {
            return false
        }

        if memo[n] != -1 {
            return memo[n] == 1
        }

        for i := 1; i*i <= n; i++ {
            square := i*i
            if !dp(n - square) {
                memo[n] = 1
                return true
            }
        }

        memo[n] = 0
        return false
    }

    return dp(n)
}

// dp (top-down + memoization), math
// time: O(nsqrt(n))
// space: O(n)