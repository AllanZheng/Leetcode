#1689. Partitioning Into Minimum Number Of Deci-Binary Numbers
import math
class Solution:
    def minPartitions(self, n: str) -> int:
        return max(set(n))
        #return max(int(i) for i in n)
print(Solution.minPartitions(Solution,"27346209830709182346"))