class Solution:
    def maxLength(self, nums: List[int]) -> int:
        ans = 0
        for i in range(len(nums)):
            prod = 1
            g = 0
            l = 1
            for j in range(i, len(nums)):
                prod *= nums[j]
                g = gcd(g, nums[j])
                l = lcm(l, nums[j])
                if prod == g * l:
                    ans = max(ans, j - i + 1)
        return ans

def gcd(a: int, b: int) -> int:
    while b != 0:
        a, b = b, a % b
    return a

def lcm(a: int, b: int) -> int:
    return 0 if a == 0 or b == 0 else abs((a // gcd(a, b)) * b)

# array, math
# time: O(n^2)
# space: O(1)