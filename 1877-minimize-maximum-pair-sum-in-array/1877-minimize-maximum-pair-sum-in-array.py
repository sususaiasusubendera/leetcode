class Solution:
    def minPairSum(self, nums: List[int]) -> int:
        if len(nums) == 2:
            return nums[0] + nums[1]

        nums.sort()

        ans = 0
        left, right = 0, len(nums) - 1
        while left < right:
            ans = max(ans, nums[left] + nums[right])
            left += 1
            right -= 1
        
        return ans

# array, greedy, sorting, two pointers
# time: O(nlog(n))
# space: O(1)