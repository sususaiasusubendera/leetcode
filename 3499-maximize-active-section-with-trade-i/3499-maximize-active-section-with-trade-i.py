class Solution:
    def maxActiveSectionsAfterTrade(self, s: str) -> int:
        count1 = 0
        for i in range(len(s)):
            if s[i] == '1':
                count1 += 1
        
        zero_blocks = []
        idx = 0
        while idx < len(s):
            start = idx
            while idx < len(s) and s[idx] == s[start]:
                idx += 1
            if s[start] == '0':
                zero_blocks.append(idx - start)
        
        if len(zero_blocks) < 2:
            return count1
        
        best = 0
        for i in range(1, len(zero_blocks), 1):
            best = max(best, zero_blocks[i] + zero_blocks[i-1])
        
        return best + count1

# greedy
# time: O(n)
# space: O(n)