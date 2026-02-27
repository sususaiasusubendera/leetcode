class Solution:
    def hasAllCodes(self, s: str, k: int) -> bool:
        h = set()
        for i in range(0, len(s) - k + 1):
            h.add(s[i:i+k])
        
        return len(h) == 2 ** k

# bit manipulation, hash map, string
# time: O(nk)
# space: O(2^k)