class Solution:
    def sequentialDigits(self, low: int, high: int) -> List[int]:
        seq_digits = []
        base = 1
        one = 1
        for i in range(2, 10, 1):
            one = (one * 10) + 1
            base += one
            n = base
            for j in range(10 - i):
                seq_digits.append(n)
                n += one
        
        ans = []
        for sd in seq_digits:
            if low <= sd <= high:
                ans.append(sd)
        
        return ans

# enumeration
# time: O(k)
# space: O(k)