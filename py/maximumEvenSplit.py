class Solution:
    def maximumEvenSplit(self, finalSum: int) -> List[int]:
        res = []
        if finalSum % 2 > 0:
            return res
        i = 2
        while i <= finalSum:
            res.append(i)
            finalSum -= i
            i += 2
        res[-1] += finalSum
        return res