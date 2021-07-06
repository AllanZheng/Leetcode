#1855. Maximum Distance Between a Pair of Values
import math
from typing import List
class Solution:
    def maxDistance(self, nums1: List[int], nums2: List[int]) -> int:
        ma = 0
        for i in range(len(nums1)):
            low = i 
            high = len(nums2)-1
            while low<high:
                mid = low +int((high-low)/2)
                
                if nums2[mid]>=nums1[i]:
                    low = mid + 1
                    ma = max(ma,mid-i)
                else:
                    high = mid     
            if nums2[high]>=nums1[i]:   
                ma = max(ma,high-i)
        return ma
print(Solution.maxDistance(Solution,[55,30,5,4,2],[100,20,10,10,5]))
print(Solution.maxDistance(Solution, [2,2,2],[10,10,1]))
print(Solution.maxDistance(Solution, [30,29,19,5],[25,25,25,25,25]))
print(Solution.maxDistance(Solution, [9,8,7],[3,2,1]))