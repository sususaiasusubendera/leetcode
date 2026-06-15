class Solution:
    def totalWaviness(self, num1: int, num2: int) -> int:
        ans = 0
        for num in range(num1, num2 + 1):
            ans += solve(num)
        
        return ans

def solve(n):
    s = str(n)

    if len(s) < 3:
        return 0
    
    waviness = 0
    for i in range(1, len(s) - 1):
        if s[i - 1] < s[i] > s[i + 1] or s[i - 1] > s[i] < s[i + 1]:
            waviness += 1
    
    return waviness

# brute force
# time: O((num2 - num1) * log(num2))
# space: O(log(num2))