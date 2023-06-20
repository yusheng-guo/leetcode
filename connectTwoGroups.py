class Solution:
    def connectTwoGroups(self, cost: List[List[int]]) -> int:
        size1, size2 = len(cost), len(cost[0])
        m = 1 << size2
        dp1, dp2 = [float("inf")] * m, [0] * m
        dp1[0] = 0

        for i in range(1, size1 + 1):
            for s in range(m):
                dp2[s] = float("inf")
                for k in range(size2):
                    if (s & (1 << k)) == 0:
                        continue
                    dp2[s] = min(dp2[s], dp2[s ^ (1 << k)] + cost[i - 1][k])
                    dp2[s] = min(dp2[s], dp1[s] + cost[i - 1][k])
                    dp2[s] = min(dp2[s], dp1[s ^ (1 << k)] + cost[i - 1][k])
            dp1 = dp2[:]
        
        return dp1[m - 1]