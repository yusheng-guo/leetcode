class Solution:
    def vowelStrings(self, words: List[str], queries: List[List[int]]) -> List[int]:
        def isVowelString(word):
            return isVowelLetter(word[0]) and isVowelLetter(word[-1])

        def isVowelLetter(c):
            return c == 'a' or c == 'e' or c == 'i' or c == 'o' or c == 'u'

        n = len(words)
        prefix_sums = [0] * (n + 1)
        for i in range(n):
            value = 1 if isVowelString(words[i]) else 0
            prefix_sums[i + 1] = prefix_sums[i] + value
        ans = []
        for i in range(len(queries)):
            start, end = queries[i]
            ans.append(prefix_sums[end + 1] - prefix_sums[start])
        return ans
