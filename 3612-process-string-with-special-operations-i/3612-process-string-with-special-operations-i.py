class Solution:
    def processStr(self, s: str) -> str:
        result = []
        for c in s:
            if c == "*":
                if result:
                    result.pop()
            elif c == "#":
                result += result
            elif c == "%":
                result.reverse()
            else:
                result.append(c)

        return "".join(result)

# string
# time: O(nk)
# space: O(n)