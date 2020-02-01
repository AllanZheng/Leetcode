from typing import List
class Solution:
    def minTimeToVisitAllPoints(self, points: List[List[int]]) -> int:
        res = 0
        for i in range(len(points)):
            if i < len(points)-1:
                x = abs(points[i][0]-points[i+1][0])
                y = abs(points[i][1]-points[i+1][1])
                if x>=y:
                    res = res + x
                else:
                    res = res + y
        return res

print(Solution.minTimeToVisitAllPoints(Solution, [[3,2],[-2,2]]))
'''
discuss里面的一行写法，思路是一样的
Python 3

    def minTimeToVisitAllPoints(self, points: List[List[int]]) -> int:
        return sum(max(abs(p[i][0] - p[i - 1][0]), abs(p[i][1] - p[i - 1][1])) for i in range(1, len(p)))
or better readability at the cost of O(n) space:

    def minTimeToVisitAllPoints(self, points: List[List[int]]) -> int:
        return sum(max(abs(cur[0] - prev[0]), abs(cur[1] - prev[1])) for cur, prev in zip(points, points[1 :]))
or use zip once again to make it clearer - credit to @infmount:

    def minTimeToVisitAllPoints(self, points: List[List[int]]) -> int:
        return sum(max(abs(c - p) for c, p in zip(cur, prev)) for cur, prev in zip(points, points[1 :]))
'''



