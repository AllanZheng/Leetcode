#1277. Count Square Submatrices with All Ones
class Solution(object):
    def countSquares(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: int
        """
        sumres = 0
        res = [[0 for i in range(len(matrix[0]))]for j in range(len(matrix))]
        for i in range(0,len(matrix)):
            for j in range(0,len(matrix[0])):
                if matrix[i][j] == 1:
                    if i==0 or j==0 :
                        res[i][j] = 1

                    else:

                        res[i][j] = 1+ min(res[i-1][j],res[i][j-1],res[i-1][j-1])
                sumres += res[i][j]

        return sumres

print(Solution.countSquares(Solution,[
    [1, 0, 1],
    [1, 1, 0],
    [1, 1, 0]
]))