from typing import List
class Solution:
    def furthestBuilding(self, heights: List[int], bricks: int, ladders: int) -> int:
        rec = []
        for i in range(len(heights)-1):
            if heights[i+1]>heights[i]:
                if len(rec)<ladders:
                    rec.append(heights[i+1]-heights[i])
                    # if rec[0]>heights[i+1]-heights[i]:
                    #     mid = rec[len(rec)-1]
                    #     rec[len(rec)-1] = rec[0]
                    #     rec[0] = mid
                    rec.sort()
                else:

                        if ladders ==0 or heights[i+1]-heights[i]<=rec[0]:
                            min = heights[i+1]-heights[i]
                        else:
                            min = rec[0]
                            rec[0]= heights[i+1]-heights[i]
                            rec.sort()
                        if min <=bricks:
                            bricks  = bricks - min
                        else:
                            return i


        return len(heights)-1
print(Solution.furthestBuilding(Solution,[1,2,3,4],1,1))


