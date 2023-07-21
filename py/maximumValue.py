class Solution:
    def maximumValue(self, strs: List[str]) -> int:
        res = 0
        for s in strs:
            is_digits = all(c.isdigit() for c in s)
            res = max(res, int(s) if is_digits else len(s))
        return res