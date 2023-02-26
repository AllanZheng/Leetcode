#287. Find the Duplicate Number
from typing import List
class Solution:
    def findDuplicate(self, nums: List[int]) -> int:
      low = 1
      high = len(nums) - 1
      while (low < high) :
        print(low,high)
        mid = int(low + (high - low) * 0.5)
        cnt = 0
        for a in nums :
            if a <= mid:
                cnt= cnt+1
        
        if cnt <= mid:
            low = mid + 1
        else:
             high = mid         
      return low
print(Solution.findDuplicate(Solution,[1,2,3,3,4,5,6,7,8]))

# 利用抽屉原理，题目明确数字范围只在[1，n]且只重复一次，如果重复在[low,mid]则重复次数超过mid,否则则在[mid+1,high]