class Solution:
    def sortPermutation(self, nums: List[int]) -> int:
        res = -1
        for i, num in enumerate(nums):
            if i != num:
                res &= num
        return 0 if res == -1 else res
    
# hint 1
# array, bit manipulation
# time: O(n)
# space: O(1)