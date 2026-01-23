class Solution:
    def minimumPairRemoval(self, nums: List[int]) -> int:
        if len(nums) == 1:
            return 0
        
        ops = 0
        while len(nums) > 1:
            sorted = True
            min_sum = 1_000_000_007
            target_idx = -1
            for i in range(len(nums) - 1):
                if nums[i + 1] < nums[i]:
                    sorted = False
                
                sum = nums[i] + nums[i + 1]
                if sum < min_sum:
                    min_sum = sum
                    target_idx = i
            
            if sorted:
                break
            
            ops += 1
            nums[target_idx] = min_sum
            nums = nums[:target_idx+1] + nums[target_idx+2:]

        return ops

# array, simulation
# time: O(n^2)
# space: O(n)