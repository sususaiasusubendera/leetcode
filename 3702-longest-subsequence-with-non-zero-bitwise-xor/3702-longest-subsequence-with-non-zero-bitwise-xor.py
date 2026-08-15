class Solution:
    def longestSubsequence(self, nums: List[int]) -> int:
        total_xor = 0
        all_zero = True
        for num in nums:
            total_xor ^= num
            if num > 0:
                all_zero = False
        
        if total_xor:
            return len(nums)
        
        if all_zero:
            return 0
        
        return len(nums) - 1

# array, bit manipulation
# time: O(n)
# space: O(1)