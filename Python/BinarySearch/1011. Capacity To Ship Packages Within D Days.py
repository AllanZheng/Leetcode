from typing import List
class Solution:
    def shipWithinDays(self, weights: List[int], D: int) -> int:
        minimum, maximum = max(weights), sum(weights)
        while minimum < maximum:
            mid = int((maximum + minimum) / 2)
            cur = 0
            days = 0
            for i in weights:
                if cur + i <= mid:
                    cur = cur + i
                else:
                    days += 1
                    cur = i
            if days < D:
                maximum = mid
            else:
                minimum = mid + 1
        return minimum

print(Solution.shipWithinDays(Solution, [1,2,3,4,5,6,7,8,9,10],5))