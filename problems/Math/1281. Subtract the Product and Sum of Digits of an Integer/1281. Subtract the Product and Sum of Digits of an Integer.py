class Solution:
    def subtractProductAndSum(self, n: int) -> int:
        res = 0
        sumofdigits = 0
        product = 1
        curstr = str(n)
        for i in curstr:
            cur = int(i, 10)

            product = product * cur
            sumofdigits = sumofdigits + cur
        res = product - sumofdigits
        return  res
print(Solution.subtractProductAndSum(Solution,234))
'''
其他优秀解法：
Python:
Space O(logN)

    def subtractProductAndSum(self, n):
        A = map(int, str(n))
        return reduce(operator.mul, A) - sum(A)
'''