class Solution:
    def minMirrorPairDistance(self, nums: List[int]) -> int:
        prev = dict()
        ans = 1_000_000_007
        for i, num in enumerate(nums):
            if num in prev:
                ans = min(ans, i - prev[num])
            prev[int(str(num)[::-1])] = i  # prev[reverse(num)]
        return -1 if ans == 1_000_000_007 else ans

# array, hash map
# time: O(n)
# space: O(n)