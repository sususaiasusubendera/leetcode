class Solution:
    def minOperations(self, nums: list[int], x: int) -> int:
        total = 0
        for num in nums:
            total += num
        
        target = total - x
        if target == 0:
            return len(nums)

        longest = 0
        window_sum = 0
        left = 0
        for right in range(len(nums)):
            window_sum += nums[right]
            while left < len(nums) and window_sum > target:
                window_sum -= nums[left]
                left += 1
            if window_sum == target:
                longest = max(longest, right - left + 1)
        
        return -1 if longest == 0 else len(nums) - longest

# array, sliding window
# time: O(n)
# space: O(1)