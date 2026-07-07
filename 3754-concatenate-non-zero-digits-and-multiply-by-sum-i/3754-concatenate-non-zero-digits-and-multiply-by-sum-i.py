class Solution:
    def sumAndMultiply(self, n: int) -> int:
        s = str(n)
        res = []
        sum = 0
        for i in range(len(s)):
            if s[i] == '0':
                continue
            res.append(s[i])
            sum += int(s[i])
        
        if len(res) == 0: return 0
        
        return int(''.join(res)) * sum

# math
# time: O(n)
# space: O(n)