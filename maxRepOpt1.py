class Solution:
    def maxRepOpt1(self, text: str) -> int:
        count = Counter(text)
        res = 0
        for i in range(len(text)):
            # step1: 找出当前连续的一段 [i, j)
            j = i
            while j < len(text) and text[j] == text[i]:
                j += 1

            # step2: 如果这一段长度小于该字符出现的总数，并且前面或后面有空位，则使用 cur_cnt + 1 更新答案
            cur_cnt = j - i
            if cur_cnt < count[text[i]] and (j < len(text) or i > 0):
                res = max(res, cur_cnt + 1)

            # step3: 找到这一段后面与之相隔一个不同字符的另一段 [j + 1, k)，如果不存在则 k = j + 1
            k = j + 1
            while k < len(text) and text[k] == text[i]:
                k += 1
            res = max(res, min(k - i, count[text[i]]))
            i = j
        return res