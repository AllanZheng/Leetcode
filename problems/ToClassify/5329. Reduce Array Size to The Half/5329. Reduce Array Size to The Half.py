from typing import List
class Solution:
    def minSetSize(self, arr: List[int]) -> int:
        res = 0
        dic = {}
        for i in arr :
            if dic.get(i) == None:
               dic[i] = 1
            else:
                dic[i] += 1

        dic = sorted(dic.items(), key=lambda d: d[1],reverse=True)
        dic = dict(dic)
        cur  = 0
        for i in dic.keys():
            res = res + 1
            cur = cur + dic[i]

            if cur >= len(arr)/2.0:
                return res
        return res
print(Solution.minSetSize(Solution, [3,3,3,3,5,5,5,2,2,7]))