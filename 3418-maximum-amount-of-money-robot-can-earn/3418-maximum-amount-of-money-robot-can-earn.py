class Solution:
    def maximumAmount(self, coins: List[List[int]]) -> int:
        m, n = len(coins), len(coins[0])
        INF = int(1e18)

        dp = [[[-INF] * 3 for _ in range(n)] for _ in range(m)]

        # base
        if coins[0][0] >= 0:
            for k in range(3):
                dp[0][0][k] = coins[0][0]
        else:
            dp[0][0][0] = coins[0][0]  # not using neutralize
            dp[0][0][1] = 0
            dp[0][0][2] = 0

        for i in range(m):
            for j in range(n):
                if i == 0 and j == 0:
                    continue
                for k in range(3):
                    best = -INF

                    if i > 0:
                        # using neutralize
                        best = max(best, dp[i-1][j][k] + coins[i][j])
                        # not using neutralize
                        if coins[i][j] < 0 and k > 0:
                            best = max(best, dp[i-1][j][k-1] + 0)

                    if j > 0:
                        # using neutralize
                        best = max(best, dp[i][j-1][k] + coins[i][j])
                        # not using neutralize
                        if coins[i][j] < 0 and k > 0:
                            best = max(best, dp[i][j-1][k-1] + 0)

                    dp[i][j][k] = best
                    
        return max(dp[m-1][n-1][0], dp[m-1][n-1][1], dp[m-1][n-1][2])

# array, dp (bot-up), matrix
# time: O(mn)
# space: O(mn)
