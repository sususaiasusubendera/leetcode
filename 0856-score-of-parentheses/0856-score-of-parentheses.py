class Solution:
    def scoreOfParentheses(self, s: str) -> int:
        def solve(left, right):
            if left + 1 == right:
                return 1
            
            depth = 0
            for i in range(left, right + 1):
                if s[i] == '(':
                    depth += 1
                elif s[i] == ')':
                    depth -= 1
                
                if not depth:
                    if i == right: # (A)
                        return 2 * solve(left + 1, right - 1)
                    else: # AB
                        return solve(left, i) + solve(i + 1, right)
            
        return solve(0, len(s) - 1)

# divide and conquer, recursive, string
# time: O(n^2)
# space: O(n)
