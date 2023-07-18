class Solution:
    def tilingRectangle(self, n: int, m: int) -> int:
        def dfs(x: int, y: int, cnt: int) -> None:
            nonlocal ans
            if cnt >= ans:
                return
            if x >= n:
                ans = cnt
                return
            
            # 检测下一行
            if y >= m:
                dfs(x + 1, 0, cnt)
                return
            # 如当前已经被覆盖，则直接尝试下一个位置
            if rect[x][y]:
                dfs(x, y + 1, cnt)
                return
            
            k = min(n - x, m - y)
            while k >= 1 and isAvailable(x, y, k):
                fillUp(x, y, k, True)
                # 跳过 k 个位置开始检测
                dfs(x, y + k, cnt + 1)
                fillUp(x, y, k, False)
                k -= 1

        def isAvailable(x: int, y: int, k: int) -> bool:
            for i in range(k):
                for j in range(k):
                    if rect[x + i][y + j] == True:
                        return False
            return True
        
        def fillUp(x: int, y: int, k: int, val: bool) -> None:
            for i in range(k):
                for j in range(k):
                    rect[x + i][y + j] = val

        ans = max(n, m)
        rect = [[False] * m for _ in range(n)]
        dfs(0, 0, 0)
        return ans