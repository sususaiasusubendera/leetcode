class Solution:
    def longestMonotonicSubarray(self, nums: List[int]) -> int:
        ans = 1

        # inreasing
        for i in range(len(nums)):
            for j in range(i, len(nums)):
                if i != j:
                    if nums[j] <= nums[j - 1]:
                        break
                    ans = max(ans, j - i + 1)
        
        # decreasing
        for i in range(len(nums)):
            for j in range(i, len(nums)):
                if i != j:
                    if nums[j] >= nums[j - 1]:
                        break
                    ans = max(ans, j - i + 1)
        
        return ans

# array
# time: O(n^2)
# space: O(1)