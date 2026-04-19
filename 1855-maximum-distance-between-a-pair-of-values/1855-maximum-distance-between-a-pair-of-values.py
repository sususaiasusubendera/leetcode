class Solution:
    def maxDistance(self, A: List[int], B: List[int]) -> int:
        i, j = 0, 1

        while i < len(A) and j < len(B):
            i += A[i] > B[j]
            j += 1

        return j - i - 1

# solution from solutions (la_castille)
# notice me senpai