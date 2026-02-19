class Solution:
    def countBinarySubstrings(self, s: str) -> int:
        groups = [1]
        i = 0
        for i in range(1, len(s)):
            if s[i-1] != s[i]:
                groups.append(1)
                continue
            groups[-1] += 1
        
        ans = 0
        for i in range(1, len(groups)):
            ans += min(groups[i-1], groups[i])
        
        return ans


# string
# time: O(n)
# space: O(1)