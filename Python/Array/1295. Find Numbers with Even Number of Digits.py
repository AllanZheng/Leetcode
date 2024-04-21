class Solution:
    def findNumbers(self, nums) -> int:
        res  = 0
        for i in nums:
            cur = str(i)
            if len(cur)%2 ==0:
                res = res + 1
        return res
print(Solution.findNumbers(Solution,[12,345,2,6,7896]))
'''
Discuss
我是用转字符串方法，其他方法可以取对数或者取模。

'''
