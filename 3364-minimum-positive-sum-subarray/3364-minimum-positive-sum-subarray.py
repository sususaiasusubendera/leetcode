class Solution:
    def minimumSumSubarray(self, nums: List[int], l: int, r: int) -> int:
        n = len(nums)

        pre = [0] * (n + 1)
        for i in range(n):
            pre[i + 1] = pre[i] + nums[i]
        
        ans = 999999999
        for size in range(l, r + 1):
            for i in range(size, n + 1):
                sum = pre[i] - pre[i - size]
                if sum > 0 and sum < ans:
                    ans = sum
        
        if ans == 999999999:
            return -1

        return ans

# array, prefix sum
# time: O(n^2)
# space: O(n)