class Solution(object):
    def smallerNumbersThanCurrent(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        res = []
        for i in range(len(nums)):
            cur = 0
            for j in range(len(nums)):

                if  nums[i]>nums[j] and i != j:
                     cur = cur + 1
                
            res.append(cur)
        return res

print(Solution.smallerNumbersThanCurrent(Solution,[8,1,2,2,3]))
