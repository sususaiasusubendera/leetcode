class Solution:
    def shortestBeautifulSubstring(self, s: str, k: int) -> str:
        ss = ""
        left = count = 0
        for right in range(len(s)):
            if s[right] == '1':
                count += 1
            
            if count == k:
                while s[left] == '0':
                    left += 1
                
                curr = s[left:right+1]
                if ss == '' or len(curr) < len(ss) or (len(curr) == len(ss) and curr < ss):
                    ss = curr

                left += 1
                count -= 1
        
        return ss

# sliding window, string
# time: O(n)
# space: O(1)