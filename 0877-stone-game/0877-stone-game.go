func stoneGame(piles []int) bool {
    n := len(piles)

    memo := make([][]int, n) // memo[l][r] stores the result of dp(l, r)
    visited := make([][]bool, n)

    for i := 0; i < n; i++ {
        memo[i] = make([]int, n)
        visited[i] = make([]bool, n)
    }

    var dp func(l, r int) int
    dp = func(l, r int) int {
        if l == r {
            return piles[l]
        }

        if visited[l][r] {
            return memo[l][r]
        }

        left := piles[l] - dp(l+1, r)
        right := piles[r] - dp(l, r-1)

        memo[l][r] = max(left, right)
        visited[l][r] = true

        return memo[l][r]
    }

    return dp(0, n-1) >= 0
}

// dp (top-down) + memoization
// time: O(n^2)
// space: O(n^2)