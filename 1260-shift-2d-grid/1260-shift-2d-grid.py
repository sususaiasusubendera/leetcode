class Solution:
    def shiftGrid(self, grid: List[List[int]], k: int) -> List[List[int]]:
        r, c = len(grid), len(grid[0])
        n = r * c
        k = k % n

        if not k: return grid

        rev(grid, 0, n - 1)
        rev(grid, 0, k - 1)
        rev(grid, k, n - 1)

        return grid

def rev(grid, i, j):
    c = len(grid[0])
    while i < j:
        grid[i // c][i % c], grid[j // c][j % c] = grid[j // c][j % c], grid[i // c][i % c]
        i += 1
        j -= 1

# array, simulation
# time: O(n)
# space: O(1)

# this approach is (prolly) the most optimal one