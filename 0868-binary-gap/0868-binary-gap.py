class Solution:
    def binaryGap(self, n: int) -> int:
        pos = []
        idx = 0
        while n > 0:
            if n & 1 == 1:
                pos.append(idx)
            idx += 1
            n >>= 1
        
        if len(pos) == 1: return 0

        ans = 0
        for i in range(1, len(pos)):
            ans = max(ans, pos[i] - pos[i - 1])
        
        return ans

# bit manipulation
# time: O(log(n))
# space: O(log(n))