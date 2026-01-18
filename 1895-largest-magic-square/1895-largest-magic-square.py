class Solution:
    def largestMagicSquare(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])

        # prefix sum horizontal
        pre_h = [[0]*(n+1) for _ in range(m+1)]
        for i in range(m):
            for j in range(n):
                pre_h[i+1][j+1] = pre_h[i+1][j] + grid[i][j]

        # prefix sum vertical
        pre_v = [[0]*(n+1) for _ in range(m+1)]
        for i in range(m):
            for j in range(n):
                pre_v[i+1][j+1] = pre_v[i][j+1] + grid[i][j]

        ans = 1
        for i in range(m):
            for j in range(n):
                idx_i, idx_j = i-1, j-1
                while idx_i >= 0 and idx_j >= 0:
                    sum = pre_h[idx_i+1][j+1] - pre_h[idx_i+1][idx_j]
                    magic = True

                    # horizontal
                    for k in range(idx_i, i+1):
                        curr_sum = pre_h[k+1][j+1] - pre_h[k+1][idx_j]
                        if curr_sum != sum:
                            magic = False
                            break

                    # vertical
                    if magic:
                        for k in range(idx_j, j+1):
                            curr_sum = pre_v[i+1][k+1] - pre_v[idx_i][k+1]
                            if curr_sum != sum:
                                magic = False
                                break
                    
                    # diagonal 1
                    if magic:
                        ii, ij = idx_i, idx_j
                        curr_sum = 0
                        while ii <= i and ij <= j:
                            curr_sum += grid[ii][ij]
                            ii += 1
                            ij += 1
                        if curr_sum != sum:
                            magic = False
                    
                    # diagonal 2
                    if magic:
                        ii, ij = i, idx_j
                        curr_sum = 0
                        while ii >= idx_i and ij <= j:
                            curr_sum += grid[ii][ij]
                            ii -= 1
                            ij += 1
                        if curr_sum != sum:
                            magic = False

                    if magic:
                        side = i - idx_i + 1
                        ans = max(ans, side) 

                    idx_i -= 1
                    idx_j -= 1

        return ans

# array, matrix, prefix sum
# time: O(mn * (min(m, n))^2)
# space: O(mn)