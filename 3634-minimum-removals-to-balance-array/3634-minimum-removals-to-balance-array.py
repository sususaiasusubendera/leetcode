class Solution:
    def minRemoval(self, nums: List[int], k: int) -> int:
        if len(nums) == 1: return 0

        nums.sort()

        left = 0
        max_window = 1
        for right in range(1, len(nums)):
            while nums[right] > nums[left] * k:
                left += 1
            
            if nums[right] <= nums[left] * k:
                max_window = max(max_window, right - left + 1)
        
        return len(nums) - max_window

# array, sliding window, sorting
# time: O(nlog(n))
# space: O(1)