class Solution:
    def minimumDeletions(self, nums: List[int]) -> int:
        if len(nums) == 1:
            return 1

        n = len(nums)
        
        minIdx = maxIdx = -1
        minVal, maxVal = 100000, -100000
        for i in range(n):
            if nums[i] < minVal:
                minVal = nums[i]
                minIdx = i
            
            if nums[i] > maxVal:
                maxVal = nums[i]
                maxIdx = i
            
        left = min(minIdx, maxIdx)
        right = max(minIdx, maxIdx)

        return min(right + 1, n - left, left + 1 + n - right)

# array, greedy
# time: O(n)
# space: O(1)