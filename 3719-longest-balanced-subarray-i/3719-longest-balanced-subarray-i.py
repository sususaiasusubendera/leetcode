class Solution:
    def longestBalanced(self, nums: List[int]) -> int:
        ans = 0
        for i in range(0, len(nums)):
            even, odd = set(), set()
            for j in range(i, len(nums)):
                if nums[j] % 2 == 0:
                    even.add(nums[j])
                else:
                    odd.add(nums[j])
                
                if len(even) == len(odd): ans = max(ans, j - i + 1)
        
        return ans

# array, hash map
# time: O(n^2)
# space: O(n)