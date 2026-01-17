class Solution:
    def largestSquareArea(self, bottomLeft: List[List[int]], topRight: List[List[int]]) -> int:
        n = len(bottomLeft)

        ans = 0
        for i in range(0, n - 1):
            for j in range(i + 1, n):
                w = min(topRight[i][0], topRight[j][0]) - max(bottomLeft[i][0], bottomLeft[j][0])
                h = min(topRight[i][1], topRight[j][1]) - max(bottomLeft[i][1], bottomLeft[j][1])

                if w > 0 and h > 0:
                    side = min(w, h)
                    area = side * side
                    ans = max(ans, area)
        
        return ans

# array, geometry, math
# time: O(n^2)
# space: O(1)