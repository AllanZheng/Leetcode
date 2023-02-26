#1395. Count Number of Teams
class Solution(object):
    def numTeams(self, rating):
        """
        :type rating: List[int]
        :rtype: int
        """
        res = 0
        dicbig = {}
        dicsmall = {}
        for i in range(len(rating)):
            bigger = []
            smaller = []
            for j in range(i+1,len(rating)):
                if rating[i]<rating[j]:
                    bigger.append(rating[j])
                if rating[i]>rating[j]:
                    smaller.append(rating[j])
            dicbig[rating[i]] = bigger
            dicsmall[rating[i]] = smaller

        for i in rating:
            for j in dicbig[i]:
                res = res + len(dicbig[j])
            for j in dicsmall[i]:
                res = res + len(dicsmall[j])
        return res

print(Solution.numTeams(Solution,[1,2,3,4]))
