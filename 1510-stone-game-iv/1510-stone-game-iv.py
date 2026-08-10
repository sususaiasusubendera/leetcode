class Solution:
    def winnerSquareGame(self, n: int) -> bool:
        memo = [-1] * (n + 1) # -1: not visited, 0: lose, 1: win
        for i in range(len(memo)):
            memo[i] = -1

        def dp(n: int) -> bool:
            if n == 0: return False

            if memo[n] != -1:
                return memo[n] == 1
            
            i = 1
            while i * i <= n:
                square = i * i
                if not dp(n - square):
                    memo[n] = 1
                    return True
                i += 1

            memo[n] = 0
            return False
        
        return dp(n)

# dp (top-down + memoization), math
# time: O(nsqrt(n))
# space: O(n)