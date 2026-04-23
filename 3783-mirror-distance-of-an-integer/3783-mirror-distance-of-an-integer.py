class Solution:
    def mirrorDistance(self, n: int) -> int:
        return abs(n - int(str(n)[::-1]))

# math
# time: O(d)
# space: O(d)