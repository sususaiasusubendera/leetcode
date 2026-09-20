class Solution:
    def totalNumbers(self, digits: List[int]) -> int:
        # 3-digit even numbers
        seen = [False] * 1000
        ans = 0
        for i in range(len(digits)):
            if digits[i] == 0:
                continue
            for j in range(len(digits)):
                if j == i:
                    continue
                for k in range(len(digits)):
                    if k == i or k == j or digits[k] % 2 != 0:
                        continue
                    num = digits[i] * 100 + digits[j] * 10 + digits[k]
                    if not seen[num]:
                        seen[num] = True
                        ans += 1
        return ans

# array, brute force
# time: O(n^3)
# space: O(1)