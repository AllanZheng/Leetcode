class Solution:
    def defangIPaddr(self, address: str) -> str:
        res = ""
        for i in range(len(address)):

            if address[i] == ".":
                res = res + "[.]"
            else:
                res = res + address[i]

        return res

print(Solution.defangIPaddr(Solution,"1.1.1.1"))

