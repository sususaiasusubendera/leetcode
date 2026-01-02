class Solution:
    def repeatedNTimes(self, nums: List[int]) -> int:
        target = len(nums) / 2
        freq = {}
        for num in nums:
            freq[num] = freq.get(num, 0) + 1
            if freq[num] == target:
                return num
        
        return -1

# array, hash table
# time: O(n)
# space: O(n)