class Solution:
    def freqAlphabets(self, s: str) -> str:
        alp = "abcdefghijklmnopqrstuvwxyz"
        res = ""
        dict= {}
        for i in range(1,27):
            cur = str(i)
            if i >= 10:
                cur = cur + "#"
            dict[cur] = alp[i-1]

        i = 0
        while i < (len(s)):
            cur = s[i]
            if i+2 < len(s):
                if s[i+2] =="#":
                    cur = cur + s[i+1]+s[i+2]
                    i = i + 2
            res = res + dict[cur]

            i = i + 1

        return res

print(Solution.freqAlphabets(Solution,"12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#"))





