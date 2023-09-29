package main

import "fmt"

func gcdOfStrings(str1 string, str2 string) string {
	byte1 := []byte(str1)
	byte2 := []byte(str2)

	commondiviors := []int{}
	min := 0
	len1 := len(byte1)
	len2 := len(byte2)
	if len1 < len2 {
		min = len1
	} else {
		min = len2
	}
	for i := min; i >= 2; i-- {
		if len1%i == 0 && len2%i == 0 {
			commondiviors = append(commondiviors, i)
		}
	}
	commondiviors = append(commondiviors, 1)
	for i := 0; i < len(commondiviors); i++ {
		rep1 := len1 / commondiviors[i]
		rep2 := len2 / commondiviors[i]
		judge := true
		res := []byte{}
		for j := 0; j < commondiviors[i]; j++ {
			for a := 1; a < rep1; a++ {
				if byte1[j] != byte1[j+a*commondiviors[i]] {
					judge = false
					break
				}
			}
			for b := 1; b < rep2; b++ {
				if byte2[j] != byte2[j+b*commondiviors[i]] {
					judge = false
					break
				}
			}
			if byte1[j] != byte2[j] {
				judge = false
				break
			}
			res = append(res, byte1[j])
		}
		if judge {
			result := string(res)
			return result
		}
	}
	return ""
}

func main() {
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))
	fmt.Println(gcdOfStrings("LEET", "CODE"))
	fmt.Println(gcdOfStrings("ABAB", "AB"))
	fmt.Println(gcdOfStrings("AAAAAA", "AAA"))
	fmt.Println(gcdOfStrings("A", "C"))
	fmt.Println(gcdOfStrings("ABABAB", "ABAB"))

}
