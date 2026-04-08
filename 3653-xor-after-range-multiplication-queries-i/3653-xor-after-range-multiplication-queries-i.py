class Solution:
    def xorAfterQueries(self, nums: List[int], queries: List[List[int]]) -> int:
        MOD = 1_000_000_007

        for q in queries:
            l, r, k, v = q[0], q[1], q[2], q[3]
            while l <= r:
                nums[l] = (nums[l] * v) % MOD
                l += k
        
        ans = 0
        for n in nums:
            ans ^= n
        
        return ans

# array, brute force
# time: O(qn)
# space: O(1)