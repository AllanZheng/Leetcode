#1340. Jump Game V
class Solution:
    def maxJumps(self, A, d):
        n = len(A)
        res = [0] * n

        def dp(i):
            if res[i]: return res[i]
            res[i] = 1
            for di in [-1, 1]:
                for j in range(i + di, i + d * di + di, di):
                    if not (0 <= j < n and A[j] < A[i]): break
                    res[i] = max(res[i], dp(j) + 1)
            return res[i]
        for i in range(n):
            print(dp(i))
        return max(map(dp, range(n)))
print(Solution.maxJumps(Solution,[7,6,5,4,3,2,1],1))
