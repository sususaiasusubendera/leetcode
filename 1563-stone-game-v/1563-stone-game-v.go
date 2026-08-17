func stoneGameV(stoneValue []int) int {
    n := len(stoneValue)

    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, n)
    }

    var dp func(left, right int) int
    dp = func(left, right int) int {
        if left == right {
            return 0
        }

        if memo[left][right] != 0 {
            return memo[left][right]
        }

        sum := 0
        for i := left; i <= right; i++ {
            sum += stoneValue[i]
        }

        sumLeft := 0
        for i := left; i < right; i++ {
            sumLeft += stoneValue[i]
            sumRight := sum - sumLeft
            if sumLeft < sumRight {
                val := dp(left, i) + sumLeft
                if val > memo[left][right] {
                    memo[left][right] = val
                }
            } else if sumLeft > sumRight {
                val := dp(i+1, right) + sumRight
                if val > memo[left][right] {
                    memo[left][right] = val
                }
            } else {
                val := max(dp(left, i), dp(i+1, right)) + sumLeft
                if val > memo[left][right] {
                    memo[left][right] = val
                }
            }
        }

        return memo[left][right]
    }

    return dp(0, n-1)
}

// array, dp top-down + memoization
// time: O(n^3)
// space: O(n^2)