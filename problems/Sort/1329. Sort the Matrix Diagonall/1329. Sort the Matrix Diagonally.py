from typing import List
class Solution:
    def diagonalSort(self, mat: List[List[int]]) -> List[List[int]]:
        for i in range(len(mat)):
            mat = Solution.diagonaltraversal(Solution,mat,i,0)
        for j in range(1,len(mat[0])):
            mat = Solution.diagonaltraversal(Solution,mat, 0, j)
        return mat

    def diagonaltraversal(self,mat:List[List[int]],row:int,column:int)-> List[List[int]]:
        cur = []
        currow = row
        curcolumn = column
        while currow<len(mat) and curcolumn<len(mat[0]):
            cur.append(mat[currow][curcolumn])
            currow += 1
            curcolumn += 1
        cur.sort()
        currow = row
        curcolumn = column
        for i in cur:
            mat[currow][curcolumn]=i
            currow += 1
            curcolumn += 1
        return mat

    def diagonalSort2(self, mat: List[List[int]]) -> List[List[int]]:
        rec = []
        res = [[0 for i in range(len(mat[0]))] for i in range(len(mat))]
        for i in mat:
            for j in i:
                rec.append(j)
        rec.sort()

        row = len(mat) - 1
        column = 0
        loop = column
        for i in rec:
            res[row][column] = i
            if row > column:
                row = row - 1
                continue
            elif column <len(mat[0])-1:
                column = column + 1
                continue
            loop = loop + 1
            column = loop
            row = len(mat)-1
        return  res
print(Solution.diagonalSort(Solution,[[3,3,1,1],[2,2,1,2],[1,1,1,2]]))
