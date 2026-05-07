class Solution:
    def rotateTheBox(self, boxGrid: List[List[str]]) -> List[List[str]]:
        m, n = len(boxGrid), len(boxGrid[0])
        for i in range(m):
            left, right = n - 1, n - 1

            while right >= 0 and boxGrid[i][right] != ".":
                right -= 1
                left -= 1

            while left >= 0 and right >= 0:
                if boxGrid[i][left] == "*":
                    right = left - 1
                elif boxGrid[i][left] == "#":
                    boxGrid[i][left] = "."
                    boxGrid[i][right] = "#"
                    right -= 1
                left -= 1
        
        newBox = [["." for _ in range(m)] for _ in range(n)]
        for i in range(n):
            for j in range(m):
                newBox[i][j] = boxGrid[m - j - 1][i]
        
        return newBox

# array, matrix, two pointers
# time: O(mn)
# space: O(nm)