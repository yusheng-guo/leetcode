class Solution:
    def minFallingPathSum(self, matrix: List[List[int]]) -> int:
        dp = [matrix[0]]
        n = len(matrix)
        for i in range(1, n):
            cur = [0] * n
            for j in range(n):
                mn = dp[-1][j]
                if j > 0:
                    mn = min(mn, dp[-1][j - 1])
                if j < n - 1:
                    mn = min(mn, dp[-1][j + 1])
                cur[j] = mn + matrix[i][j]
            dp.append(cur)
        return min(dp[-1])