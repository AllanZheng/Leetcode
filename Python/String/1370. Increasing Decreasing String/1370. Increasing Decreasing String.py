#1370. Increasing Decreasing String

class Solution(object):
    def sortString(self, s):
        """
        :type s: str
        :rtype: str
        """
        alp = "abcdefghijklmnopqrstuvwxyz"
        recalp = "zyxwvutsrqponmlkjihgfedcba"
        res = ""
        dict = {}
        for i in range(26):
            dict[alp[i]] = 0
        for i in s:
            dict[i] += 1
        rec = len(s)

        while rec >0:
            for i in alp:
                if dict[i]>0:
                    res = res + i
                    dict[i] -=1
                    rec -= 1
            for i in recalp:
                if dict[i]>0:
                    res = res + i
                    dict[i] -=1
                    rec -=1
        return res

print(Solution.sortString(Solution, "cccccc"))
