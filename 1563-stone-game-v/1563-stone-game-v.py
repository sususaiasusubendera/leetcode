class Solution:
    def stoneGameV(self, stoneValue: List[int]) -> int:
        @cache
        def dp(left: int, right: int) -> int:
            if left == right:
                return 0
            
            sum_total = 0
            for i in range(left, right + 1):
                sum_total += stoneValue[i]
            
            ans = 0
            sum_left = 0
            for i in range(left, right):
                sum_left += stoneValue[i]
                sum_right = sum_total - sum_left
                if sum_left < sum_right:
                    ans = max(ans, dp(left, i) + sum_left)
                elif sum_left > sum_right:
                    ans = max(ans, dp(i + 1, right) + sum_right)
                else:
                    ans = max(ans, max(dp(left, i), dp(i + 1, right)) + sum_left)
            
            return ans
        
        return dp(0, len(stoneValue) - 1)

# dp top-down + memoization (@cache)  
# time: O(n^3)
# space: O(n^2)