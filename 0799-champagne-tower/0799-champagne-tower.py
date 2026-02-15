class Solution:
    def champagneTower(self, poured: int, query_row: int, query_glass: int) -> float:
        dp =[[0] * k for k in range(1, 102)]
        dp[0][0] = poured
        for r in range(query_row + 1): # row
            for g in range(r + 1): # glass
                e = (dp[r][g] - 1.0) / 2.0
                if e > 0:
                    dp[r + 1][g] += e
                    dp[r + 1][g + 1] += e
        
        return min(1, dp[query_row][query_glass])

# dp
# time: O(query_row^2)
# space: O(1)