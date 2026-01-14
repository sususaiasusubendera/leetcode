func minimumDeleteSum(s1 string, s2 string) int {
    n, m := len(s1), len(s2)

	dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, m+1)
    }

    // base case
    // s1 = "" -> delete s2
    for j := m - 1; j >= 0; j-- {
        dp[n][j] = dp[n][j+1] + int(s2[j])
    }
    // s2 = "" -> delete s1
    for i := n - 1; i >= 0; i-- {
        dp[i][m] = dp[i+1][m] + int(s1[i])
    }

    for i := n - 1; i >= 0; i-- {
        for j := m - 1; j >= 0; j-- {
            if s1[i] == s2[j] { // skip
                dp[i][j] = dp[i+1][j+1]
                continue
            }

            delS1 := int(s1[i]) + dp[i+1][j]
            delS2 := int(s2[j]) + dp[i][j+1]
            if delS1 < delS2 {
                dp[i][j] = delS1
            } else {
                dp[i][j] = delS2
            }
        }
    }

    return dp[0][0]
}

// dp
// time: O(nm)
// space: O(nm)

// visit me again senpai