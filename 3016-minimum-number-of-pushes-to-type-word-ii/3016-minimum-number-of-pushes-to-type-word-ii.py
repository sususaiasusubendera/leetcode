class Solution:
    def minimumPushes(self, word: str) -> int:
        freq = {}
        for i in range(len(word)):
            freq[word[i]] = freq.get(word[i], 0) + 1
        
        letters = []
        for l in freq:
            letters.append(l)
        
        letters.sort(key=lambda c: freq[c], reverse=True)

        keypads = [0] * 26
        count = 0
        ans = 0
        for i in range(len(letters)):
            c = letters[i]
            if not keypads[ord(c) - ord('a')]:
                keypads[ord(c) - ord('a')] = (count // 8) + 1
                ans += keypads[ord(c) - ord('a')] * freq[c]
                count += 1
            else:
                ans += keypads[ord(c) - ord('a')] * freq[c]
        
        return ans

# greedy, hash map, sorting, string
# time: O(n + klog(k))
# space: O(k)