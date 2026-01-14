class Solution:
    def maxDotProduct(self, nums1: List[int], nums2: List[int]) -> int:
        NEG_INF = -10**10
        n, m = len(nums1), len(nums2)

        dp = [[NEG_INF for _ in range(m+1)] for _ in range (n+1)]

        for i in range(n - 1, -1, -1):
            for j in range(m - 1, - 1, -1):
                take = nums1[i] * nums2[j]
                if dp[i+1][j+1] > 0:
                    take += dp[i+1][j+1]
                
                dp[i][j] = max(take, max(dp[i][j+1], dp[i+1][j]))

        return dp[0][0]

# array, dp
# time: O(nm)
# space: O(nm)