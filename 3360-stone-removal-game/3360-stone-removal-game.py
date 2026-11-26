class Solution:
    def canAliceWin(self, n: int) -> bool:
        count = 0
        remove = 10
        while n >= remove:
            n -= remove
            remove -= 1
            count += 1
        return False if count % 2 == 0 else True

# math
# time: O(sqrt(n))
# space: O(1)