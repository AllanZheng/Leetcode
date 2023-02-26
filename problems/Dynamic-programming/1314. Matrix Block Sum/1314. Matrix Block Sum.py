# 1314. Matrix Block Sum
class Solution(object):
    def matrixBlockSum(self, mat, K):
        """
        :type mat: List[List[int]]
        :type K: int
        :rtype: List[List[int]]
        """
        res = [[0 for i in range(len(mat[0]))]for i in range(len(mat))]
        for i in range(len(res)):
            for j in range(len(res[0])):
                left,right = Solution.find(self,i,K,0,len(res)-1)

                left1, right1 = Solution.find(self, j, K, 0, len(res)-1)
                for row in range(left,right+1,1):
                    for col in range(left1,right1+1,1):

                        res[i][j] = res[i][j] + mat[row][col]
        return res

    def find(self,cur,K,leftbound,rightbound):
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
