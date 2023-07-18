class Solution:
    def f(self, s: str) -> int:
        cnt = 0
        ch = 'z'
        for c in s:
            if c < ch:
                ch = c
                cnt = 1
            elif c == ch:
                cnt += 1
        return cnt

    def numSmallerByFrequency(self, queries: List[str], words: List[str]) -> List[int]:
        count = [0] * 12
        for s in words:
            count[self.f(s)] += 1
        for i in range(9, 0, -1):
            count[i] += count[i + 1]
        res = []
        for s in queries:
            res.append(count[self.f(s) + 1])
        return res