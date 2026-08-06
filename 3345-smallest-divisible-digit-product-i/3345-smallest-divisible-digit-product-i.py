class Solution:
    def smallestNumber(self, n: int, t: int) -> int:
        for i in range(n, 101):
            if prod_digits(i) % t == 0:
                return i
        
        return -1

def prod_digits(n: int) -> int:
    res = 1
    while n > 0:
        res *= n % 10
        n //= 10
    
    return res

# math
# time: O(1)
# space: O(1)