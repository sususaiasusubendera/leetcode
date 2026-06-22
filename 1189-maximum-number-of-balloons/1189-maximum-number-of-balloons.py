class Solution:
    def maxNumberOfBalloons(self, text: str) -> int:
        freq = {}
        for c in text:
            if c not in freq:
                freq[c] = 0
            freq[c] += 1
        
        b, a, l, o, n = freq.get('b', 0), freq.get('a', 0), freq.get('l', 0), freq.get('o', 0), freq.get('n', 0)
        
        count = 0
        while b > 0 and a > 0 and l > 1 and o > 1 and n > 0:
            count += 1
            b -= 1
            a -= 1
            l -= 2
            o -= 2
            n -= 1

        return count

# hash map, string
# time: O(n)
# space: O(1) 