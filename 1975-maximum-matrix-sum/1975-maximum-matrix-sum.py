class Solution:
    def maxMatrixSum(self, matrix: List[List[int]]) -> int:
        sum = 0
        count_neg = 0
        zero_exists = False
        min_pos, min_neg = 100_000, 100_000 # from constraint
        for i in range(len(matrix)):
            for j in range(len(matrix[0])):
                sum += abs(matrix[i][j])
                if matrix[i][j] < 0:
                    count_neg += 1
                    min_neg = min(min_neg, -matrix[i][j])
                elif matrix[i][j] == 0:
                    zero_exists = True
                else:
                    min_pos = min(min_pos, matrix[i][j])

        if count_neg % 2 == 0 or zero_exists:
            return sum
        
        return sum - (2 * min(min_pos, min_neg))

# array, greedy, matrix
# time: O(n^2)
# space: O(1)