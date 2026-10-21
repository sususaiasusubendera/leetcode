class Solution:
    def minOperations(self, s: str) -> int:
        min_char, max_char = min(s), max(s)

        # sorted
        if all(a <= b for a, b in zip(s, s[1:])):
            return 0

        if len(s) == 2:
            return -1
        
        if s[0] == min_char or s[-1] == max_char:
            return 1

        # min_char and max_char in the middle
        for i in range(1, len(s) - 1):
            if s[i] == min_char or s[i] == max_char:
                return 2
        
        # s[0] == max_char and s[len(s) - 1] == min_char
        return 3

# string
# time: O(n)
# space: O(1)