from typing import List
class Solution:
    def sumZero(self, n: int) -> List[int]:
        res = []
        cur = 0
        last = 0
        for i in range(n):
            if i!=n-1:
                cur = cur - 1
                last = last - cur
                res.append(cur)

            else:
                res.append(last)
        return res
print(Solution.sumZero(Solution,5))

'''
disucss 一行代码写法
return list(range(1,n)) + [-(n-1)*n/2, ]
'''