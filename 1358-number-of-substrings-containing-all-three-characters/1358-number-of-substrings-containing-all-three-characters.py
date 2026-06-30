class Solution:
    def numberOfSubstrings(self, s: str) -> int:
        freq = {}
        ans = 0
        left = 0
        for right in range(len(s)):
            freq[s[right]] = freq.get(s[right], 0) + 1
            while freq.get('a', 0) > 0 and freq.get('b', 0) > 0 and freq.get('c', 0) > 0:
                ans += len(s) - right
                freq[s[left]] -= 1
                left += 1
        
        return ans

# hash map, sliding window, string
# time: O(n)
# space: O(1)