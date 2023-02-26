class Solution:
    def countVowelStrings(self, n: int) -> int:
        dict = {}
        for i in range(n):
            if i == 0:
                dict[1]=1
                dict[2]=1
                dict[3]=1
                dict[4]=1
                dict[5]=1
            else:
                dict[5] = dict[5]+dict[4]+dict[3]+dict[2]+dict[1]
                dict[4]  = dict[4]+dict[3]+dict[2]+dict[1]
                dict[3] =  dict[3]+dict[2]+dict[1]
                dict[2] =  dict[2]+dict[1]
                dict[1]= dict[1]
        return dict[5]+dict[4]+dict[3]+dict[2]+dict[1]
print(Solution.countVowelStrings(Solution,33))


