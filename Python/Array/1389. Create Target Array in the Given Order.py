#1389. Create Target Array in the Given Order
class Solution(object):
    def createTargetArray(self, nums, index):
        """
        :type nums: List[int]
        :type index: List[int]
        :rtype: List[int]
        """
        res = []
        for i in range(len(nums)):
            if index[i]>=len(res):
                res.append(nums[i])
            else:
                res = res[:index[i]] + nums[i:i+1] + res[index[i]:]


        return res


print(Solution.createTargetArray(Solution,[1,2,3,4,0],[0,1,2,3,0]))