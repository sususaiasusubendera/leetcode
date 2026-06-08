class Solution:
    def leftRightDifference(self, nums: List[int]) -> List[int]:
        leftSum, rightSum = 0, 0
        for i in range(len(nums) - 1, 0, -1):
            rightSum += nums[i]
        
        ans = [0] * len(nums)
        for i in range(len(nums)):
            ans[i] = abs(leftSum - rightSum)

            leftSum += nums[i]
            if i < len(nums) - 1:
                rightSum -= nums[i + 1]

        return ans

# array, prefix sum
# time: O(n)
# space: O(1)