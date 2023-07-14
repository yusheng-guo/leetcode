class Solution:
    move = 0

    def distributeCoins(self, root: Optional[TreeNode]) -> int:
        def dfs(root):
            moveleft = 0
            moveright = 0
            if root is None:
                return 0
            if root.left is not None:
                moveleft = dfs(root.left)
            if root.right is not None:
                moveright = dfs(root.right)
            self.move += abs(moveleft) + abs(moveright)
            return moveleft + moveright + root.val - 1

        dfs(root)
        return self.move