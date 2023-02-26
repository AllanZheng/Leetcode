class solution:

    def decompressRLElist(self, nums):
        res = []
        for i in range(len(nums)):
            if i % 2 == 0:
               for j in range(nums[i]):
                   res.append(nums[i+1])
        return res


print(solution.decompressRLElist(solution,[1,2,3,4]))
