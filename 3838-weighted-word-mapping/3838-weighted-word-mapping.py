class Solution:
    def mapWordWeights(self, words: List[str], weights: List[int]) -> str:
        ans = []
        for word in words:
            weight = 0
            for i in range(len(word)):
                weight += weights[ord(word[i]) - ord('a')]
        
            ans.append(chr(ord('z') - (weight % 26)))
        
        return ''.join(ans)

# array, string
# time: O(nk)
# space: O(n)