from typing import List
class Solution:
    def oddCells(self, n: int, m: int, indices: List[List[int]]) -> int:
        martix = [[0 for i in range(m)] for i in range(n)]
        res = 0
        for i in indices:
                for row in range(len(martix)):
                    for column in range(len(martix[0])):
                        if column == i[1]:
                            martix[row][column]= martix[row][column]+1
                        if row == i[0]:
                            martix[row][column]= martix[row][column]+1
        for row in range(len(martix)):
            for column in range(len(martix[0])):
                if martix[row][column]%2!=0:
                    res = res + 1
        return res
print(Solution.oddCells(Solution,2,2,[[1,1],[0,0]]))


