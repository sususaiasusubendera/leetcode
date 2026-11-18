class Solution:
    def minChanges(self, s: str) -> int:
        count = 0
        for i in range(0, len(s) // 2):
            if s[(i * 2)] != s[(i * 2) + 1]:
                count += 1
        return count

# string
# time: O(n/2)
# space: O(1)