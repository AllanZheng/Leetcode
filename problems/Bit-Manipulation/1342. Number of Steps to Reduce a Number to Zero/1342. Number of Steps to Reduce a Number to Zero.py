# 1342. Number of Steps to Reduce a Number to Zero
class Solution(object):
    def numberOfSteps (self, num):
        """
        :type num: int
        :rtype: int
        """
        step = 0
        while num != 0:
            if num%2 != 0:
                num = num -1
            else:
                num = num/2
            step = step + 1
        return step
'''
正确解法
去数数字的二进制标签表示里面有多少位0和1，0直接加一只做除法，1减一再做除2，当遇到最后一个1的时候
int numberOfSteps (int num) {
		if(!num) return 0;
        int res = 0;
        while(num) {
            res += (num & 1) ? 2 : 1;
            num >>= 1;
        }
        return res - 1;
    }
'''