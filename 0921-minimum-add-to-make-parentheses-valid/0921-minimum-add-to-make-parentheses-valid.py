class Solution:
    def minAddToMakeValid(self, s: str) -> int:
        stack = []
        for c in s:
            if stack and stack[len(stack) - 1] == '(':
                if c == ')':
                    stack.pop()
                    continue
            stack.append(c)
        
        return len(stack)

# greedy, stack, string
# time: O(n)
# space: O(n)