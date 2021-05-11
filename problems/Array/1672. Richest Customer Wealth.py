from typing import List
class Solution:
    def maximumWealth(self, accounts: List[List[int]]) -> int:
        maxwealth = 0
        for i in accounts:
            cur = 0 
            for j in i:
                cur = cur + j
            if cur>maxwealth:
                maxwealth = cur
        return maxwealth

print(Solution.maximumWealth(Solution,[[2,8,7],[7,1,3],[1,9,5]]))