# 1314. Matrix Block Sum
class Solution(object):
    def matrixBlockSum(self, mat, K):
        """
        :type mat: List[List[int]]
        :type K: int
        :rtype: List[List[int]]
        """
        res = [[0 for i in range(len(mat[0]))]for i in range(len(mat))]
        rang = [[0 for i in range(len(mat[0])+1)]for i in range(len(mat)+1)]
        for i in range(1,len(res)+1):
            for j in range(1,len(res[0])+1):

                    rang[i][j] = rang[i-1][j] + rang[i][j-1] + mat[i-1][j-1]-rang[i-1][j-1]
        print(rang)
        for i in range(len(res)):
            for j in range(len(res[0])):
                left, right = Solution.find(self, i, K, 0, len(res)-1)

                left1, right1 = Solution.find(self, j, K, 0, len(res[0])-1)

                res[i][j] = rang[right+1][right1+1] - rang[left][right1+1] - rang[right+1][left1] + rang[left][left1]
        return res
    def find(self, cur, K, leftbound, rightbound):
            if cur - K >= leftbound:
                left = cur - K
            else:
                left = leftbound
            if cur + K <= rightbound:
                right = cur + K
            else:
                right = rightbound
            return left,right




print(Solution.matrixBlockSum(Solution,[[1,2,3],[4,5,6],[7,8,9]],2))
