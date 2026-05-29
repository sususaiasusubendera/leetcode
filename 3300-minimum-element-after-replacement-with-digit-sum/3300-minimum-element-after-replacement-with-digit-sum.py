class Solution:
    def minElement(self, nums: List[int]) -> int:
        a = [0] * len(nums)
        for i in range(len(nums)):
            a[i] = countSumDigit(nums[i])

        ans = 50
        for num in a:
            ans = min(ans, num)

        return ans


def countSumDigit(n):
    sum = 0
    while n > 0:
        sum += n % 10
        n //= 10
    return sum

# array
# time: O(nd)
# space: O(1)