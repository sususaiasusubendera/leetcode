class Solution:
    def validSequence(self, word1: str, word2: str) -> List[int]:
        n, m = len(word1), len(word2)

        last = [-1] * m
        
        j = m - 1
        for i in range(n - 1, -1, -1):
            if j >= 0 and word1[i] == word2[j]:
                last[j] = i
                j -= 1
        
        ans = []
        skip = 0
        j = 0
        for i in range(n):
            if j == m: # optimization
                break
            
            if word1[i] == word2[j] or (skip == 0 and (j == m - 1 or i < last[j + 1])):
                if word1[i] != word2[j]:
                    skip += 1
                ans.append(i)
                j += 1
        
        if j == m:
            return ans
        
        return []

# greedy (very tricky), string, two pointers
# time: O(n + m)
# space: O(m)