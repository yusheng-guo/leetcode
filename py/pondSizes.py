class Solution:
    def pondSizes(self, land: List[List[int]]) -> List[int]:
        m, n = len(land), len(land[0])
        
        def bfs(x: int, y: int) -> int:
            res = 0
            q = deque([(x, y)])
            land[x][y] = -1

            while q:
                x, y = q.popleft()
                res += 1
                for dx in [-1, 0, 1]:
                    for dy in [-1, 0, 1]:
                        if dx == dy == 0:
                            continue
                        if x + dx < 0 or x + dx >= m or y + dy < 0 or y + dy >= n or land[x + dx][y + dy] != 0:
                            continue
                        land[x + dx][y + dy] = -1
                        q.append((x + dx, y + dy))
            return res
        
        res = list()
        for i in range(m):
            for j in range(n):
                if land[i][j] == 0:
                    res.append(bfs(i, j))
        res.sort()
        return res