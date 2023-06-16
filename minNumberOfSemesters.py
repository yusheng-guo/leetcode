class Solution:
    def minNumberOfSemesters(self, n: int, relations: List[List[int]], k: int) -> int:
        dp = [inf] * (1 << n)
        need = [0] * (1 << n)
        for edge in relations:
            need[(1 << (edge[1] - 1))] |= 1 << (edge[0] - 1)
        dp[0] = 0
        for i in range(1, (1 << n)):
            need[i] = need[i & (i - 1)] | need[i & (-i)]
            if (need[i] | i) != i: # i 中有任意一门课程的前置课程没有完成学习
                continue
            sub = valid = i ^ need[i] # 当前学期可以进行学习的课程集合
            if sub.bit_count() <= k: # 如果个数小于 k，则可以全部学习，不再枚举子集
                dp[i] = min(dp[i], dp[i ^ sub] + 1)
            else: # 枚举子集
                while sub:
                    if sub.bit_count() <= k:
                        dp[i] = min(dp[i], dp[i ^ sub] + 1)
                    sub = (sub - 1) & valid
        return dp[-1]