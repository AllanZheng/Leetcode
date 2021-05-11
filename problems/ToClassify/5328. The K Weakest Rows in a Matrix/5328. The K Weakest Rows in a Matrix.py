from typing import List
class Solution:
    def kWeakestRows(self, mat: List[List[int]], k: int) -> List[int]:
        dic = {}
        for row in range(len(mat)):
            cur = 0
            for column in mat[row]:
                if column == 1:
                    cur = cur + 1
            dic[row] = cur
        dic = sorted(dic.items(), key=lambda d: d[1])
        dic = dict(dic)
        res = []
        for i in dic.keys():
            res.append(i)
            if len(res)==k:
                return  res
        return  res
print(Solution.kWeakestRows(Solution,[[1,1,0,0,0],
 [1,1,1,1,0],
 [1,0,0,0,0],
 [1,1,0,0,0],
 [1,1,1,1,1]],3))
