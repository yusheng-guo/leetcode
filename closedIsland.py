class UnionFind:
    def __init__(self, n: int) -> None:
        self.parent = list(range(n))
        self.rank = [0] * n
    
    def uni(self, x: int, y: int) -> None:
        parent_ = self.parent
        rank_ = self.rank

        rootx, rooty = self.find(x), self.find(y)
        if rootx != rooty:
            if rank_[rootx] > rank_[rooty]:
                parent_[rooty] = rootx
            elif rank_[rootx] < rank_[rooty]:
                parent_[rootx] = rooty
            else:
                parent_[rooty] = rootx
                rank_[rootx] += 1

    def find(self, x: int) -> int:
        parent_ = self.parent

        if parent_[x] != x:
            parent_[x] = self.find(parent_[x])
        return parent_[x]

class Solution:
    def closedIsland(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        uf = UnionFind(m * n)

        for i in range(m):
            if grid[i][0] == 0:
                uf.uni(i * n, 0)
            if grid[i][n - 1] == 0:
                uf.uni(i * n + n - 1, 0)
        
        for j in range(1, n - 1):
            if grid[0][j] == 0:
                uf.uni(j, 0)
            if grid[m - 1][j] == 0:
                uf.uni((m - 1) * n + j, 0)
        
        for i in range(m):
            for j in range(n):
                if grid[i][j] == 0:
                    if i > 0 and grid[i - 1][j] == 0:
                        uf.uni(i * n + j, (i - 1) * n + j)
                    if j > 0 and grid[i][j - 1] == 0:
                        uf.uni(i * n + j, i * n + j - 1)

        cnt = set()
        for i in range(m):
            for j in range(n):
                if grid[i][j] == 0:
                    cnt.add(uf.find(i * n + j))
        
        total = len(cnt)
        if uf.find(0) in cnt:
            total -= 1
        return total