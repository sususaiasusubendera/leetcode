class Solution:
    def rotatedDigits(self, n: int) -> int:
        ans = 0
        for i in range(1, n + 1, 1):
            if good(i):
                ans += 1
        
        return ans
        
def good(n):
    valid = False
    while n > 0:
        d = n % 10

        if d == 3 or d == 4 or d == 7:
            return False

        if d == 2 or d == 5 or d == 6 or d == 9:
            valid = True
        
        n //= 10

    return valid

# brute force
# time: O(nlog(n))
# space: O(1)