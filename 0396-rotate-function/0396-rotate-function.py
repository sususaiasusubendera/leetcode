class Solution:
    def maxRotateFunction(self, nums: List[int]) -> int:
        n = len(nums)
        if n < 2:
            return 0

        sum = 0
        f = 0
        for i, num in enumerate(nums):
            sum += num
            f += i * num
        
        ans = f
        for i in range(1, n):
            f = (n * nums[i - 1]) - sum + f
            ans = max(ans, f)
        
        return ans

# array
# time: O(n)
# space: O(1)

# note (tips)
# grab some paper and a pen
# try to derive something from f(1) - f(0)