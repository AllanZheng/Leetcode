class Solution:
    def maxDepth(self, s: str) -> int:
        max  = 0
        cur  = 0
        for i in s:
            if i == "(":
                cur = cur  +  1
                if cur > max:
                  max = cur

            if i == ")":
                cur = cur-1

        return max
print(Solution.maxDepth(Solution,"1"))