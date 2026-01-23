class Solution:
    def minBitwiseArray(self, nums: List[int]) -> List[int]:
        ans = [0] * len(nums)
        for i in range(len(nums)):
            if nums[i] & 1 == 0:
                ans[i] = -1
                continue
            
            d = 1
            while nums[i] & d != 0:
                d <<= 1
            d >>= 1
            ans[i] = nums[i] - d
        
        return ans

# array, bit manipulation
# time: O(n)
# space: O(n)