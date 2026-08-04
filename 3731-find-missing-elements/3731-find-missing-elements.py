class Solution:
    def findMissingElements(self, nums: List[int]) -> List[int]:
        nums.sort()

        ans = []
        num = nums[0]
        idx = 0
        while idx < len(nums):
            while num != nums[idx]:
                ans.append(num)
                num += 1
            idx += 1
            num += 1

        return ans

# array, sorting
# time: O(nlog(n))
# space: O(n)