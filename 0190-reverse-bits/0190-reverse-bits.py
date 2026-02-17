class Solution:
    def reverseBits(self, n: int) -> int:
        bits = [0] * 32
        i = 0
        while n > 0:
            bits[i] = n % 2
            n //= 2
            i += 1
        
        ans = 0
        mul = 1
        for i in range(0, 31):
            mul *= 2
        for i in range(0, 32):
            ans += bits[i] * mul
            mul //= 2
        
        return ans

# bit manipulation
# time: O(1)
# space: O(1)