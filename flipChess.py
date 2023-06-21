class Solution:
    def flipChess(self, chessboard: List[str]) -> int:
        def judge(chessboard: List[List[str]], x: int, y: int, dx: int, dy: int) -> bool:
            x += dx
            y += dy
            while 0 <= x < len(chessboard) and 0 <= y < len(chessboard[0]):
                if chessboard[x][y] == "X":
                    return True
                elif chessboard[x][y] == ".":
                    return False
                x += dx
                y += dy
            return False
        
        def bfs(chessboard: List[str], px: int, py: int) -> int:
            chessboard = [list(row) for row in chessboard]
            cnt = 0
            q = deque([(px, py)])
            chessboard[px][py] = "X"

            while q:
                tx, ty = q.popleft()
                for dx in [-1, 0, 1]:
                    for dy in [-1, 0, 1]:
                        if dx == dy == 0:
                            continue
                        if judge(chessboard, tx, ty, dx, dy):
                            x, y = tx + dx, ty + dy
                            while chessboard[x][y] != "X":
                                q.append((x, y))
                                chessboard[x][y] = "X"
                                x += dx
                                y += dy
                                cnt += 1
            return cnt

        res = 0
        for i in range(len(chessboard)):
            for j in range(len(chessboard[0])):
                if chessboard[i][j] == ".":
                    res = max(res, bfs(chessboard, i, j))
        return res