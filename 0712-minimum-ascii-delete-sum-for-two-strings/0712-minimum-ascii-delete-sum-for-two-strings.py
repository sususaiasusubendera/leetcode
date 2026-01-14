class Solution:
    def minimumDeleteSum(self, s1: str, s2: str) -> int:
        n, m = len(s1), len(s2)

        dp = [[0 for _ in range(m+1)] for _ in range(n+1)]

        # base case
        # s1 = "" -> delete s2
        for j in range(m - 1, -1, -1):
            dp[n][j] = dp[n][j+1] + ord(s2[j])
        # s2 = "" -> delete s1
        for i in range(n - 1, -1, -1):
            dp[i][m] = dp[i+1][m] + ord(s1[i])

        for i in range(n - 1, -1, -1):
            for j in range(m - 1, -1, -1):
                if s1[i] == s2[j]:
                    dp[i][j] = dp[i+1][j+1]
                    continue
                
                del_s1 = ord(s1[i]) + dp[i+1][j]
                del_s2 = ord(s2[j]) + dp[i][j+1]
                if del_s1 < del_s2:
                    dp[i][j] = del_s1
                else:
                    dp[i][j] = del_s2
        
        return dp[0][0]

# dp
# time: O(nm)
# space: O(nm)