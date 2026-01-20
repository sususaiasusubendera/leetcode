class Solution:
    def minBitwiseArray(self, nums: List[int]) -> List[int]:
        ans = [0 for _ in range(len(nums))]
        for i in range(len(nums)):
            val = -1
            for j in range(1, nums[i]):
                if j | (j + 1) == nums[i]:
                    val = j
                    break
            ans[i] = val
        
        return ans

# aray, bit manipulation
# time: O(nm)
# space: O(n)