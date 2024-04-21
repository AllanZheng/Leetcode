class Solution:
    def groupThePeople(self, groupSizes):
        rec = [[0 for i in range(501)] for i in range(501)]
        res = []
        for i in range(len(groupSizes)):
            rec[groupSizes[i]][i]=1
        for i in range(501):
            cur = []
            for j in range(501):
                if rec[i][j] == 1:
                    if len(cur)==i:
                        res.append(cur)
                        cur = []
                    cur.append(j)
            if len(cur)>0:
                res.append(cur)
        return res
print(Solution.groupThePeople(Solution,[3,3,3,3,3,1,3]))

