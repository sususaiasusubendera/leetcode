class Solution:
    def minimumPushes(self, word: str) -> int:
        ans = 0
        letters = [0] * 26
        count = 0
        for i in range(len(word)):
            c = ord(word[i]) - ord('a')
            if not letters[c]:
                letters[c] = (count // 8) + 1
                ans += letters[c]
                count += 1
            else:
                ans += letters[c]
        
        return ans

# greedy, math, string
# time: O(n)
# space: O(1)