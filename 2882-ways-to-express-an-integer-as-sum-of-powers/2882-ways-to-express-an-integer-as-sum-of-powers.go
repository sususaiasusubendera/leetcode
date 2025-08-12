func numberOfWays(n int, x int) int {
    memo := map[[2]int]int{}
    var dp func(i, sum int) int
    dp = func(i, sum int) int {
        key := [2]int{i, sum}
        if v, ok := memo[key]; ok {
            return v
        }
    
        if sum == n {
            return 1
        }
        if sum > n {
            return 0
        }

        p := pow(i, x)
        if p > n-sum {
            return 0
        }

        res := (dp(i+1, sum) + dp(i+1, sum+p)) % (1e9+7)
        memo[key] = res
        return res
    }

    return dp(1, 0)
}

// dp (top-down + memoization)
// time: O(n^(1+1/x))
// space: O(n^(1+1/x))

func pow(n, p int) int {
    result := 1
    for i := 1; i <= p; i++ {
        result *= n
    }

    return result
}