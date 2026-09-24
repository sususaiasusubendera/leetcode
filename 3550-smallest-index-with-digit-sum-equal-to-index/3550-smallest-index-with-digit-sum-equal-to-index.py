class Solution:
    def smallestIndex(self, nums: List[int]) -> int:
        for i, num in enumerate(nums):
            if i == sum_digits(num):
                return i
        return -1

def sum_digits(n: int) -> int:
    sum = 0
    while n > 0:
        sum += n % 10
        n //= 10
    return sum

# array, math
# time: O(nd)
# space: O(1)