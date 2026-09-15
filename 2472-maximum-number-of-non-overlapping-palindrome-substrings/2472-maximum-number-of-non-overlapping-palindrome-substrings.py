class Solution:
    def maxPalindromes(self, s: str, k: int) -> int:
        ans = start = 0
        for r in range(k - 1, len(s)):
            # size k
            l = r - k + 1
            if l >= start and is_pal(l, r, s):
                ans += 1
                start = r + 1
                continue
            
            # size k + 1
            l = r - k
            if l >= start and is_pal(l, r, s):
                ans += 1
                start = r + 1
        
        return ans
        
def is_pal(l, r: int, s: str) -> bool:
    while l < r:
        if s[l] != s[r]:
            return False
        l += 1
        r -= 1
    
    return True

# greedy (the greedy idea is from editorial)
# time: O(nk)
# space: O(1)