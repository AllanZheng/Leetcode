#167. Two Sum II - Input array is sorted
from typing import List
class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        for i in range(len(numbers)):
            x = target -numbers[i]
            low = i
            high = len(numbers)-1
            while low<=high:
                mid = int((low+high)/2)
                if x ==numbers[mid]:
                    return [i+1,mid+1]
                elif x>numbers[mid]:
                    low = mid+1
                else:
                    high = mid-1
            return [-1,-1]
print(Solution.twoSum(Solution,[2,7,11,15],9))
print(Solution.twoSum(Solution,[2,3,4],6))
print(Solution.twoSum(Solution,[-1,0],-1))
