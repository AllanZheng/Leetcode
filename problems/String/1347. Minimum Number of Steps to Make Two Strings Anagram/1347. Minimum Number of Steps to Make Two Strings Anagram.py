# 1347. Minimum Number of Steps to Make Two Strings Anagram
class Solution(object):
    def minSteps(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: int
        """
        alp = "abcdefghijklmnopqrstuvwxyz"
        res = 0
        sdict = {}
        tdict = {}
        for i in range(26):
            sdict[alp[i]] = 0
            tdict[alp[i]] = 0
        for i in t:
            tdict[i] = tdict[i]+1
        for i in s:
            if sdict[i] == tdict[i]:
                res = res + 1
            else:
                sdict[i] = sdict[i] + 1
        return res
print(Solution.minSteps(Solution,"friend","family"))

