class Solution:
    def maxTotalValue(self, nums: List[int], k: int) -> int:
        min_num, max_num = float("inf"), 0
        for num in nums:
            min_num = min(min_num, num)
            max_num = max(max_num, num)
        
        return k * (max_num - min_num)

# array, greedy
# time: O(n)
# space: O(1)