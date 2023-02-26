package main

import "fmt"

func getSmallestString(n int, k int) string {
	alp := map[int]string{
		1:"a",
		2:"b",
		3:"c",
		4:"d",
		5:"e",
		6:"f",
		7:"g",
		8:"h",
		9:"i",
		10:"j",
		11:"k",
		12:"l",
		13:"m",
		14:"n",
		15:"o",
		16:"p",
		17:"q",
		18:"r",
		19:"s",
		20:"t",
		21:"u",
		22:"v",
		23:"w",
		24:"x",
		25:"y",
		26:"z",
	}
	mo := k%26
	re := k/26
	res :=""
	if n-re>=mo{
		x := (n-re-mo)/26
		//xmo :=(n-re)%26
		for i:=0;i<n-re-1+x;i++{
			res+="a"
		}
		fmt.Println(26-(k-(n-re-1+x)-(re-x)*26))
		res+=alp[26-(n-re-mo-x*26)]
		for i:=0;i<re-x;i++{
			res +="z"
		}
	}else{


	for i:=0;i<n-re-1;i++{
		res += "a"
	}
	res+=alp[mo-(n-re-1)]
	for i:=0;i<re;i++{
		res +="z"
	}
	}
	fmt.Println(len(res))
	return res
}

func main(){
	fmt.Println(getSmallestString(80,576))
}