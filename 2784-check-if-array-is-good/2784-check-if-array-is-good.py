class Solution:
    def isGood(self, nums: List[int]) -> bool:
        if len(nums) <= 1:
            return False

        nums.sort()

        if nums[-1] != nums[-2] and nums[-1] != len(nums) - 1:
            return False

        for i, num in enumerate(nums):
            if i != len(nums) - 1:
                if num != i + 1:
                    return False
        
        return True
        
# array, sorting
# time: O(nlog(n))
# space: O(1)