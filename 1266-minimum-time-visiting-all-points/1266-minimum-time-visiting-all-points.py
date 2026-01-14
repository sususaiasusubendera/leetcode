class Solution:
    def minTimeToVisitAllPoints(self, points: List[List[int]]) -> int:
        if len(points) == 1: return 0

        ans = 0
        for i in range(1, len(points)):
            x1, y1 = points[i-1][0], points[i-1][1]
            x2, y2 = points[i][0], points[i][1]
            
            dx = abs(x2 - x1)
            dy = abs(y2 - y1)

            ans += max(dx, dy)
        
        return ans

# array, geometry, math
# time: O(n)
# space: O(1)