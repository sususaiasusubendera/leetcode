class Solution:
    def minMoves2(self, nums: list[int]) -> int:
        med = find_median(nums)
        res = 0
        for num in nums:
            res += abs(med - num)
        return res

def find_median(nums: list[int]) -> int:
    nums.sort()

    if len(nums) % 2 == 1:
        return nums[len(nums) // 2]
    
    return (nums[len(nums) // 2] + nums[(len(nums) // 2) - 1]) // 2

# array, math, median, sorting
# time: O(nlog(n))
# space: O(1)