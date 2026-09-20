class Solution:
    def reverseDegree(self, s: str) -> int:
        map = [0] * 26
        n = 26
        for i in range(len(map)):
            map[i] = n
            n -= 1
        
        ans = 0
        for i, v in enumerate(s):
            ans += map[ord(v) - ord('a')] * (i + 1)
        
        return ans
    
    # string
    # time: O(s)
    # space: O(1)