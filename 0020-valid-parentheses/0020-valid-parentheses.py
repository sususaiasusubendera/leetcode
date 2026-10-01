class Solution:
    def isValid(self, s: str) -> bool:
        pairs = {
            "(": ")",
            "{": "}",
            "[": "]",
        }

        if len(s) % 2 == 1:
            return False

        stack = []
        for c in s:
            if c in pairs:
                stack.append(c)
            elif not len(stack) or c != pairs[stack[-1]]:
                return False
            else:
                stack.pop()
        
        return len(stack) == 0

# stack, string
# time: O(n)
# space: O(m)