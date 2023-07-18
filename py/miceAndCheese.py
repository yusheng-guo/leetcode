class Solution:
    def miceAndCheese(self, reward1: List[int], reward2: List[int], k: int) -> int:
        ans = 0
        n = len(reward1)
        pq = []
        for i in range(n):
            ans += reward2[i]
            heappush(pq, reward1[i] - reward2[i])
            if len(pq) > k:
                heappop(pq)
        while pq:
            ans += heappop(pq)
        return ans