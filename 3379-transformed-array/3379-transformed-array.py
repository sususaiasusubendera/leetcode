class Solution:
    def constructTransformedArray(self, nums: List[int]) -> List[int]:
        n = len(nums)
        result = [0] * n
        for i in range(0, n):
            result[i] = nums[(i + nums[i]) % n]
        
        return result

# array, simulation
# time: O(n)
# space: O(n)