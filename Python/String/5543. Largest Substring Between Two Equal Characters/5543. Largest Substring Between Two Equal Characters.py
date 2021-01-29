class Solution:
    def maxLengthBetweenEqualCharacters(self, s: str) -> int:
        alp = "abcdefghijklmnopqrstuvwxyz"
        dict = {}
        max = -1
        for i in range(26):
            dict[alp[i]] = 0
        for i in s:
            for j in range(26):
                if alp[j] == i:
                   if dict[i]>0:

                        if max < dict[i]-1:
                            max = dict[i]-1
                        dict[i] = dict[i] + 1
                   elif dict[i]==0:
                        dict[i]=1

                elif dict[alp[j]]>0:
                        dict[alp[j]] = dict[alp[j]]+1

        return max


print(Solution.maxLengthBetweenEqualCharacters(Solution,"dcabbaccdccccd"))