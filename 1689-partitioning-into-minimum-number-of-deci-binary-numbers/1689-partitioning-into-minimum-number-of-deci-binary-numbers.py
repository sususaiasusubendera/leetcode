class Solution:
    def minPartitions(self, n: str) -> int:
        m = 0
        for i in range(0, len(n)):
            m = max(m, int(n[i]))
        return m

# greedy, string
# time: O(n)
# space: O(1)