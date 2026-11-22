class Solution:
    def checkString(self, s: str) -> bool:
        valid = True
        if s[0] == 'b':
            valid = 'a' not in s
        else:
            b = False
            for c in s:
                if not b and c == 'b':
                    b = True
                elif b and c == 'a':
                    valid = False
                    break
        
        return valid

# string
# time: O(n)
# space: O(1)