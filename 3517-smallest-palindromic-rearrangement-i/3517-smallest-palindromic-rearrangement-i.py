class Solution:
    def smallestPalindrome(self, s: str) -> str:
        if len(s) == 1: return s

        temp = list(s)
        left, right = [], []

        if len(s) % 2 == 0:
            for i in range(len(temp) // 2):
                left.append(temp[i])
                right.append(temp[len(temp) - 1 - i])
            
            left.sort()
            right.sort(reverse=True)

            return "".join(left + right)
        
        idx = 0
        while idx < len(temp) // 2:
            left.append(temp[idx])
            right.append(temp[len(temp) - 1 - idx])
            idx += 1
        
        mid = temp[idx]
        left.sort()
        right.sort(reverse=True)

        return "".join(left + list(mid) + right)

# sorting, string
# time: O(nlog(n))
# space: O(n)