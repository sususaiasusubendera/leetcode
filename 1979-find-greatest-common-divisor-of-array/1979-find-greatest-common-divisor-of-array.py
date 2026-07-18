class Solution:
    def findGCD(self, nums: List[int]) -> int:
        s, l = 1000, 0 # smallest num, largest num
        for num in nums:
            s = min(s, num)
            l = max(l, num)
        
        return gcd(s, l)

def gcd(a, b: int) -> int:
    while b != 0:
        a, b = b, a % b
    return a

# array, math
# time: O(n)
# space: O(1)