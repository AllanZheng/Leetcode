#1828. Queries on Number of Points Inside a Circle
import math
from typing import List
class Solution:
    def countPoints(self, points: List[List[int]], queries: List[List[int]]) -> List[int]:
        res = []
        for i in queries:
            cur = 0 
            for j in points:
                if math.pow((j[0]-i[0]),2)+math.pow((j[1]-i[1]),2)-math.pow(i[2],2)<=0:
                    cur =cur+1
                    
            res.append(cur)
        return res
print(Solution.countPoints(Solution,[[1,1],[2,2],[3,3],[4,4],[5,5]],[[1,2,2],[2,2,2],[4,3,2],[4,3,3]]))