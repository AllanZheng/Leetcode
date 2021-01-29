import os
import sys
import shutil
import  string
import re

py = '.py'
go = '.go'
def file_name(file_dir):

    for root, dirs, files in os.walk(file_dir):
            if (os.path.dirname(root)== file_dir or os.path.dirname(os.path.dirname(root))== file_dir) and not root.isdigit():
                for i in files:
                    if i =='process.py' or i in dirs:
                        continue
                    if i.endswith(py) :
                            str = i.strip(py)
                            if str not in dirs:
                                filedir = root + "\\" + str
                                os.makedirs(filedir)
                            oldpath = root + "\\" + i
                            newpath = root+"\\"+str+"\\"+i
                            print(oldpath,newpath)
                            shutil.move(oldpath,newpath)
                    if i.endswith(go) :
                            str = i.strip(go)
                            if str not in dirs:
                                filedir = root + "\\" + str
                                os.makedirs(filedir)
                            oldpath = root + "\\" + i
                            newpath = root + "\\" + str + "\\" + i
                            print(oldpath, newpath)
                            shutil.move(oldpath,newpath)




path = sys.path[0]
print(path)
file_name(path)
