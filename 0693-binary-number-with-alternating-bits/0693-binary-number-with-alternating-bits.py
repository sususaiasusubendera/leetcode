class Solution:
    def hasAlternatingBits(self, n: int) -> bool:
        bit = n & 1
        altBits = True
        while n > 0:
            n >>= 1
            if bit == n & 1:
                altBits = False
                break
            bit = n & 1
        return altBits

# bit manipulation
# time: O(log(n))
# space: O(1)