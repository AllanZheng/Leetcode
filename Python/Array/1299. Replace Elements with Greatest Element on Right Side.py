from typing import List
class Solution:
    def replaceElements(self, arr: List[int]) -> List[int]:
        maximum = -1
        for i in range(len(arr) - 1, -1, -1):
            mid = maximum

            if arr[i] > maximum:
                maximum = arr[i]
                arr[i] = mid
            else:
                arr[i] = maximum
        return arr
