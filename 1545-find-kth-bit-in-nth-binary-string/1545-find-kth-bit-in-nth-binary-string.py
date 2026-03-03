class Solution:
    def findKthBit(self, n: int, k: int) -> str:
        s = "0"
        for i in range(1, n):
            s += "1" + reverse(invert(s))
        return s[k - 1]

# simulation, string
# time: O(2^n)
# space: O(2^n)

def reverse(s):
    res = list(s)
    for i in range(0, len(res) // 2):
        res[i], res[len(s) - 1 - i] = res[len(s) - 1 - i], res[i]
    return "".join(res)

def invert(s):
    res = list(s)
    for i in range(0, len(res)):
        if res[i] == "1":
            res[i] = "0"
        else:
            res[i] = "1"
    return "".join(res)