class Solution:
    def minimumCost(self, nums: List[int]) -> int:
        nums1 = nums[1:]
        nums1.sort()
        return nums[0] + nums1[0] + nums1[1]

# array, sorting
# time: O(nlog(n))
# space: O(1)