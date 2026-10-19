class Solution:
    def targetIndices(self, nums: list[int], target: int) -> list[int]:
        nums.sort()

        ans = []
        for i, num in enumerate(nums):
            if num == target:
                ans.append(i)
        
        return ans

# array, sorting
# time: O(nlog(n))
# space: O(1)