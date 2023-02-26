package main

import "strings"

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
     str1 :=""
     str2 := ""
     str1 = strings.Join(word1,"")
     str2  = strings.Join(word2,"")
     if str1==str2{
     	return true
	 }else{
	 	return false
	 }

}