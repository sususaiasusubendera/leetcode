class Solution:
    def missingInteger(self, nums: List[int]) -> int:
        s = set()
        for num in nums:
            s.add(num)

        sum = nums[0]
        for i in range(1, len(nums), 1):
            if nums[i] != nums[i - 1] + 1:
                break
            
            sum += nums[i]
        
        while sum in s:
            sum += 1
        
        return sum

# array, hash map
# time: O(n)
# space: O(n)