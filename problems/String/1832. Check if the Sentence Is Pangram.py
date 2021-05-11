#1832. Check if the Sentence Is Pangram
class Solution:
    def checkIfPangram(self, sentence: str) -> bool:
        dict = {}
        for i in sentence:
            if dict.get(i):
                dict[i]=dict[i]+1
                
            else:
               dict[i]=1
        if len(dict)!=26:
            return False
        for j in dict:
            if dict[j]>1:
                return False
        return True
print(Solution.checkIfPangram(Solution,"abcdefghijklmnopqrstuvwxyz"))