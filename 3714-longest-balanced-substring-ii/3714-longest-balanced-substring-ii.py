class Solution:
    def longestBalanced(self, s: str) -> int:
        res = [
            solve1(s, 'a'),
            solve1(s, 'b'),
            solve1(s, 'c'),
            solve2(s, 'a', 'b'),
            solve2(s, 'a', 'c'),
            solve2(s, 'b', 'c'),
            solve3(s)
        ]

        ans = 0
        for r in res:
            ans = max(ans, r)
        return ans

# hash map, prefix sum, string
# ref: tdkkdt
# time: O(n)
# space: O(n)
    
def solve1(s, t):
    res, count = 0, 0
    for i in range(0, len(s)):
        if s[i] == t:
            count += 1
            res = max(res, count)
            continue
        count = 0
    return res

def solve2(s, t1, t2):
    res, c1, c2 = 0, 0, 0
    freq = {} # diff -> freq
    for i in range(0, len(s)):
        if s[i] != t1 and s[i] != t2:
            c1, c2 = 0, 0
            freq.clear()
            continue

        if s[i] == t1:
            c1 += 1
        else:
            c2 += 1
        
        if c1 == c2:
            res = max(res, c1 + c2)
        else:
            diff = c1 - c2
            if diff in freq:
                l = i - freq[diff]
                res = max(res, l)
            else:
                freq[diff] = i
    return res
        
def solve3(s):
    res, ca, cb, cc = 0, 0, 0, 0
    freq = {}
    for i in range(0, len(s)):
        match s[i]:
            case 'a': ca += 1
            case 'b': cb += 1
            case 'c': cc += 1
        
        if ca == cb == cc:
            res = i + 1
        else:
            diff = (ca - cb, cb - cc)
            if diff in freq:
                l = i - freq[diff]
                res = max(res, l)
            else:
                freq[diff] = i
    return res