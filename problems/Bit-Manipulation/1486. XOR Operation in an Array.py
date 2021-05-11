#1486. XOR Operation in an Array
class Solution:
    def xorOperation(self, n: int, start: int) -> int:
        res = 0
        for i in range(n):
            cur = start+2*i
            res = res^cur
        return res
print(Solution.xorOperation(Solution,10,5))
