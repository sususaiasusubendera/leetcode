class Solution:
    def maxSideLength(self, mat: List[List[int]], threshold: int) -> int:
        m, n = len(mat), len(mat[0])

        # 2d prefix sum
        pre = [[0] * (n + 1) for _ in range(m + 1)]
        for i in range(m):
            for j in range(n):
                pre[i + 1][j + 1] = (
                    pre[i][j + 1] + pre[i + 1][j] - pre[i][j] + mat[i][j]
                )

        ans = 0
        sub = 0  # substractor -> top-left
        for i in range(m):
            for j in range(n):
                if i == 0 or j == 0:
                    if mat[i][j] <= threshold:
                        ans = max(ans, 1)
                    continue

                idx_i, idx_j = i - sub - 1, j - sub - 1
                temp_sub = sub
                while idx_i >= 0 and idx_j >= 0:
                    temp_sub += 1
                    sum = (
                        pre[i + 1][j + 1]
                        - pre[idx_i][j + 1]
                        - pre[i + 1][idx_j]
                        + pre[idx_i][idx_j]
                    )
                    if sum <= threshold:
                        side = i - idx_i + 1
                        ans = max(ans, side)
                        sub = max(sub, temp_sub)
                    idx_i -= 1
                    idx_j -= 1

        return ans


# array, matrix, prefix sum
# time: O(mn * min(m, n))
# space: O(mn)
