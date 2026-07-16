class Solution:
    def gcdSum(self, nums: list[int]) -> int:
        pre_gcd = [0] * len(nums)
        mx = 0
        for i in range(len(nums)):
            mx = max(mx, nums[i])
            pre_gcd[i] = gcd(nums[i], mx)
        
        pre_gcd.sort()

        ans = 0
        for i in range(len(pre_gcd) // 2):
            ans += gcd(pre_gcd[i], pre_gcd[len(pre_gcd) - 1 - i])
        
        return ans

def gcd(a, b: int) -> int:
    while b != 0:
        a, b = b, a % b
    return a

# array, math, sorting
# time: O(nlog(n))
# space: O(n)