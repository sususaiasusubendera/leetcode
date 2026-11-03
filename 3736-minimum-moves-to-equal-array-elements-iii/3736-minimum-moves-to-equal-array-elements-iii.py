class Solution:
    def minMoves(self, nums: List[int]) -> int:
        max_num = 0
        for num in nums:
            max_num = max(max_num, num)
        
        res = 0
        for num in nums:
            res += max_num - num
        
        return res

# array, math
# time: O(n)
# space: O(1)