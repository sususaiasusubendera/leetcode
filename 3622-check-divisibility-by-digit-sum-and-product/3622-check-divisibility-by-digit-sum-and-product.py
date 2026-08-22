class Solution:
    def checkDivisibility(self, n: int) -> bool:
        sum, prod = findSumProduct(n)

        return n % (sum + prod) == 0

def findSumProduct(n: int) -> tuple[int, int]:
    sum, prod = 0, 1
    while n > 0:
        d = n % 10
        sum += d
        prod *= d
        n //= 10

    return sum, prod

# math
# time: O(log(n))
# space: O(1)