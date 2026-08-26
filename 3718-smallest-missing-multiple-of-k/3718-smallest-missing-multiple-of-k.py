class Solution:
    def missingMultiple(self, nums: List[int], k: int) -> int:
        s = set()
        for num in nums:
            s.add(num)
        
        num = k
        while num < 1000:
            if num not in s:
                return num
            num += k
        
        return num

# array, hash map
# time: O(n + t)
# space: O(n)