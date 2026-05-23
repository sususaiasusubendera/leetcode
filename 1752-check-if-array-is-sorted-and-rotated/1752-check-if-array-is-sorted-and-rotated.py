class Solution:
    def check(self, nums: List[int]) -> bool:
        count_decreasing = 0
        for i in range(0, len(nums)):
            if nums[i] > nums[(i + 1) % len(nums)]:
                count_decreasing += 1
            
            if count_decreasing > 1:
                return False
        
        return True

# array
# time: O(n)
# space: O(1)