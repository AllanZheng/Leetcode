class Solution:
    def balancedStringSplit(self, s: str) -> int:
        stackr = []
        stackl = []
        res = 0
        for i in s:
            if i == "R":
                stackr.append(i)
            else:
                stackl.append(i)
            if len(stackl)== len(stackr):
                res = res + 1
                while len(stackl)!=0 and len(stackr)!=0:
                    stackr.pop()
                    stackl.pop()
        return res
print(Solution.balancedStringSplit(Solution,"RLRRRLLRLL"))


