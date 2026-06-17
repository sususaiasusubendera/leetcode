class Solution:
    def processStr(self, s: str, k: int) -> str:
        l = 0
        for c in s:
            if c == '*':
                if l > 0:
                    l -= 1
            elif c == '#':
                l *= 2
            elif c == '%':
                continue
            else:
                l += 1
        
        if k > l - 1:
            return '.'

        for i in range(len(s) - 1, -1, -1):
            if s[i] == '*':
                l += 1
            elif s[i] == '#':
                l //= 2
                if k >= l:
                    k -= l
            elif s[i] == '%':
                k = l - k - 1
            else:
                if k == l - 1:
                    return s[i]
                else:
                    l -= 1
        
        return '.'

# simulation, string
# time: O(n)
# space: O(1)