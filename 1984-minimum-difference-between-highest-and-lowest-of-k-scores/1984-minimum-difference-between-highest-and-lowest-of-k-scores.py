class Solution:
    def minimumDifference(self, nums: List[int], k: int) -> int:
        nums.sort(reverse=True)

        ans = 1_000_000_007
        left, right = 0, k - 1
        while right <= len(nums) - 1:
            ans = min(ans, nums[left] - nums[right])
            right += 1
            left += 1

        return ans
