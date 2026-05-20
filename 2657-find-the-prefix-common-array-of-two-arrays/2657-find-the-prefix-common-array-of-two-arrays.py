class Solution:
    def findThePrefixCommonArray(self, A: List[int], B: List[int]) -> List[int]:
        n = len(A)
        f = {}
        ans = [0] * n
        total = 0
        for i in range(n):
            if A[i] == B[i]:
                total += 1
            else:
                f[A[i]] = f.get(A[i], 0) + 1
                f[B[i]] = f.get(B[i], 0) + 1

                if f[A[i]] == 2:
                    total += 1
                
                if f[B[i]] == 2:
                    total += 1
        
            ans[i] += total

        return ans

# hash map
# time: O(n)
# space: O(n)

# #1 2026/05/20