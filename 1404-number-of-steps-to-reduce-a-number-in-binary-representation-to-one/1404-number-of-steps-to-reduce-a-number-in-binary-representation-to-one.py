class Solution:
    def numSteps(self, s: str) -> int:
        ops = 0
        s = list(s)
        while len(s) > 1:
            n = len(s)
            if s[n - 1] == '0':
                s.pop()
            else: # lsb == 1
                i = len(s) - 1
                while i >= 0 and s[i] != '0':
                    s[i] = '0'
                    i -= 1
                if i < 0:
                    s.insert(0, '1')
                else:
                    s[i] = '1'
            ops += 1
        
        return ops

# bit manipulation, string, simulation
# time: O(n)
# space: O(n)