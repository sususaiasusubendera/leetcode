class Solution:
    def minMoves(self, nums: list[int]) -> int:
        min_el = 9999999999
        for num in nums:
            min_el = min(min_el, num)

        res = 0
        for num in nums:
            res += num - min_el

        return res

# array,math
# time: O(n)
# space: O(1)
# a beautiful approach from guilatrova (from discussion)