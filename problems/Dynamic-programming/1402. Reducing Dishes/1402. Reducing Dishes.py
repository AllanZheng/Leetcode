#1402. Reducing Dishes
from typing import List
class Solution:
    def maxSatisfaction(self, satisfaction: List[int]) -> int:
        satisfaction.sort()
        rec = []
        for i in range(len(satisfaction)-1,0,-1):

            if i == len(satisfaction)-1:
                cur = satisfaction[i]
            else:
                cur =rec[len(rec)-1]+satisfaction[i]
            rec.append(cur)
        rec.reverse()
        choose = 0
        res = 0
        for j in range(len(satisfaction)):
            if (choose+1)*satisfaction[j]>=0 or(j<len(satisfaction)-2 and rec[j]+(choose + 1)*satisfaction[j]>0):
                res = res+ (choose+1)*satisfaction[j]
                choose += 1
        if res<0:
            res = 0
        return res



#print(Solution.maxSatisfaction(Solution, [-1,-8,0,5,-9]))
#print(Solution.maxSatisfaction(Solution, [-2,5,-1,0,3,-3]))
#print(Solution.maxSatisfaction(Solution, [4,3,2]))
#print(Solution.maxSatisfaction(Solution,[-1,-4,-5]))
print(Solution.maxSatisfaction(Solution,[2,-2,-3,1]))