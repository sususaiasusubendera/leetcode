class Solution:
    def longestBalanced(self, s: str) -> int:
        ans = 0
        for i in range(0, len(s)):
            freq = {}
            for j in range(i, len(s)):
                freq[s[j]] = freq.get(s[j], 0) + 1

                balance = True
                curr_freq = freq[s[j]]
                for _, v in freq.items():
                    if curr_freq != v:
                        balance = False
                        break
                
                if balance:
                    ans = max(ans, j - i + 1)
            
        return ans
    
# hash map, string
# time: O(n^2)
# space: O(1)