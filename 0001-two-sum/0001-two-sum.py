class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        hm = {}
        for i in range(len(nums)):
            c = target - nums[i]
            if c in hm:
                return [hm[c], i]
            hm[nums[i]] = i

# array, hash map
# time: O(n)
# space: O(n)