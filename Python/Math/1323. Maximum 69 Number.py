class Solution:
    def maximum69Number (self, num: int) -> int:
        strnum = str(num)
        for i in range(len(strnum)):
            if strnum[i] == "6":
                curlist = list(strnum)
                curlist[i] = "9"
                strnum ="".join(curlist)
                break
        num = int(strnum)
        return num
print(Solution.maximum69Number(Solution,9996))
'''
别人的简便写法，思路是一样的。
Python

str.replace has a parameter for maximum number of occurrences to replace.

def maximum69Number(self, num):
    return int(str(num).replace('6', '9', 1))
'''