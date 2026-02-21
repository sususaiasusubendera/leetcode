class Solution:
    def countPrimeSetBits(self, left: int, right: int) -> int:
        ans = 0
        for n in range(left, right + 1):
            if isPrime(countSetBits(n)): ans += 1
        return ans

# bit manipulation, brian kernighan’s algorithm
# time: O(nlog(n))
# space: O(1)

def countSetBits(n):
    count = 0
    while n > 0:
        n = n & (n - 1)
        count += 1
    return count

def isPrime(n):
    if n < 2: 
        return False

    i = 2
    while i * i <= n:
        if n % i == 0: return False
        i += 1
    return True
        