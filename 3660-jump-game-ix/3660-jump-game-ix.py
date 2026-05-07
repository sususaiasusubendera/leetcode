class Solution:
    def maxValue(self, nums: List[int]) -> List[int]:
        stack = []
        for i in range(len(nums)):
            val, left, right = nums[i], i, i
            while stack and stack[-1][0] > nums[i]:
                top_val, top_left, top_right = stack.pop()
                
                val = max(val, top_val)
                left = min(left, top_left)
                right = max(right, top_right)

            stack.append((val, left, right))

        ans = [0] * len(nums)
        for val, left, right in stack:
            for j in range(left, right + 1):
                ans[j] = val
        
        return ans

# monotonic stack
# time: O(n)
# space: O(n)