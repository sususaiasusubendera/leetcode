class Solution:
    def isTrionic(self, nums: List[int]) -> bool:
        p, q = False, False
        for i in range(1, len(nums)-1):
            if not p and not q and nums[i] <= nums[i - 1]:
                break
            
            if not p and not q and nums[i - 1] < nums[i] > nums[i + 1]:
                p = True
            
            if p and not q and nums[i - 1] > nums[i] < nums[i + 1]:
                q = True
            
            if p and q and nums[i] >= nums[i + 1]:
                p = False
                q = False
                break
        
        return p and q

# array
# time: O(n)
# space: O(1)