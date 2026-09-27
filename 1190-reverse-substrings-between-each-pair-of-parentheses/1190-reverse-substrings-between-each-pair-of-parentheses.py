class Solution:
    def reverseParentheses(self, s: str) -> str:
        open_parantheses_idx = deque()
        res = []
        for c in s:
            if c == '(':
                open_parantheses_idx.append(len(res))
            elif c == ')':
                start = open_parantheses_idx.pop()
                res[start:] = res[start:][::-1]
            else:
                res.append(c)
        return "".join(res)

# stack, string
# time: O(n^2)
# space: O(n)