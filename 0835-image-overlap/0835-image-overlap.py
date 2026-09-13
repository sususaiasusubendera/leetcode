class Solution:
    def largestOverlap(self, img1: List[List[int]], img2: List[List[int]]) -> int:
        ans = 0
        n = len(img1)
        for i in range(-(n - 1), n):
            for j in range(-(n - 1), n):
                ans = max(ans, count_one(i, j, img1, img2))
        
        return ans

def count_one(sr, sc: int, img1, img2: List[List[int]]) -> int:
    overlap = 0
    n = len(img1)
    for i in range(n):
        for j in range(n):
            ni = i + sr
            nj = j + sc
            if 0 <= ni < n and 0 <= nj < n:
                if img1[ni][nj] == 1 and img2[i][j] == 1:
                    overlap += 1
            
    return overlap

# array, brute force, matrix
# time: O(n^4)
# space: O(1)