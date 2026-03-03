class Solution:
    def concatenatedBinary(self, n: int) -> int:
        MOD = 1_000_000_007

        ans = 0
        length = 0
        for i in range(1, n + 1):
            if i & (i - 1) == 0:
                length += 1
            
            ans = ((ans << length) + i) % MOD

        return ans

# bit manipulation, math, simulation
# time: O(n)
# space: O(1)