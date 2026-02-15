class Solution:
    def addBinary(self, a: str, b: str) -> str:
        ia, ib = len(a) - 1, len(b) - 1
        res = []
        carry = False
        while ia >= 0 or ib >= 0:
            if ia >= 0 and ib >= 0 and a[ia] == '1' and b[ib] == '1':
                if carry:
                    res.append('1')
                else:
                    res.append('0')
                carry = True
            elif ia >= 0 and ib >= 0 and a[ia] == '0' and b[ib] == '0':
                if carry:
                    res.append('1')
                else:
                    res.append('0')
                carry = False
            elif ia >= 0 and ib >= 0 and ((a[ia] == '1' and b[ib] == '0') or (a[ia] == '0' and b[ib] == '1')):
                if carry:
                    res.append('0')
                    carry = True
                else:
                    res.append('1')
                    carry = False
            elif ia >= 0:
                if a[ia] == '1':
                    if carry:
                        res.append('0')
                        carry = True
                    else:
                        res.append('1')
                        carry = False
                else:
                    if carry:
                        res.append('1')
                    else:
                        res.append('0')
                    carry = False
            elif ib >= 0:
                if b[ib] == '1':
                    if carry:
                        res.append('0')
                        carry = True
                    else:
                        res.append('1')
                        carry = False
                else:
                    if carry:
                        res.append('1')
                    else:
                        res.append('0')
                    carry = False
            ia -= 1
            ib -= 1
        
        if carry: res.append('1')
        
        res.reverse()
        ans = ''.join(res)
        return ans

# bit manipulation, math, simulation, string
# time: O(n)
# space: O(n)