class Solution:
    def minimumDeletions(self, s: str) -> int:
        stack = []
        ans = 0
        for i in range(0, len(s)):
            if len(stack) == 0:
                stack.append(s[i])
                continue
            
            if stack[len(stack)-1] == 'b' and s[i] == 'a':
                ans += 1
                stack.pop()
            else:
                stack.append(s[i])
        
        return ans

# stack, string
# time: O(n)
# space: O(n)