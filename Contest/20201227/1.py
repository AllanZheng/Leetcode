class Solution:
    def halvesAreAlike(self, s: str) -> bool:
        len1 = int(len(s)/2)
        front = s[0:len1]
        end   = s[len1:]
        dict = ["a","e","i","o","u","A","E","I","O","U"]
        rec =0
        for i in front:
            if i in dict:
                rec +=1
        for j in end :
            if j in dict:
                rec -=1
        if rec == 0:
            return True
        else:
            return False

print(Solution.halvesAreAlike(Solution,"AbCdEfGh"))
