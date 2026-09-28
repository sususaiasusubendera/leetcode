class Solution:
    def evaluate(self, s: str, knowledge: List[List[str]]) -> str:
        d = dict(knowledge)
        ans, start = [], -1
        for i, c in enumerate(s):
            if c == "(":
                start = i
            elif c == ")":
                ans.append(d.get(s[start + 1 : i], "?"))
                start = -1
            elif start == -1:
                ans.append(c)
        return "".join(ans)

# array, hash map, string
# time: O(n + m)
# space: O(n + m)