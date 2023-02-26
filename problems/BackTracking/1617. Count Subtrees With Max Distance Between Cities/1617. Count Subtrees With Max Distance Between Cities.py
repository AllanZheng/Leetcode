class Solution:
    def countSubgraphsForEachDiameter(self, n: int, edges: List[List[int]]) -> List[int]:
           dict = {}
           res = []

           for  i in edges:
                dict[i[0]]= i[1]
           for  cur = 0,cur<n,cur++:
               
