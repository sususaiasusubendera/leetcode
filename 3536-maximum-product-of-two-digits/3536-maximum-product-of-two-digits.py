class Solution:
    def maxProduct(self, n: int) -> int:
        temp = []
        
        while n > 0:
            digit = n % 10
            temp.append(digit)
            n //= 10
        
        temp.sort()

        return temp[len(temp) - 1] * temp[len(temp) - 2]

# math, sorting
# time: O(klog(k))
# space: O(k)