class Solution:
    def separateDigits(self, nums: List[int]) -> List[int]:
        ans = []
        for num in nums:
            for d in str(num):
                ans.append(int(d))
        
        return ans

# array
# time: O(k)
# space: O(k)