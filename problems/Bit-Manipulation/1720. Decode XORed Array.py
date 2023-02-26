# 1720. Decode XORed Array
from typing import List
class Solution:
    def decode(self, encoded: List[int], first: int) -> List[int]:
        res =[]
        res.append(first)
        for i in range(len(encoded)):
            res.append(res[i]^encoded[i])
        return res
print(Solution.decode(Solution,[6,2,7,3],4))
