class Solution:
    def maxAlternatingSum(self, nums: List[int]) -> int:
        nums.sort(key=abs, reverse=True)

        limit = (len(nums) // 2) + 1 if len(nums) % 2 == 1 else len(nums) // 2
        res = 0
        for i in range(len(nums)):
            if i < limit:
                res += nums[i] ** 2
            else:
                res -= nums[i] ** 2
        return res

# array, sorting
# time: O(nlog(n))
# space: O(1)