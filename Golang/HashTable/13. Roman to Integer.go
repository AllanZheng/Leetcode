package main

import "fmt"

func romanToInt(s string) int {
	length := len(s)
	sb := []byte(s)
	RtoI := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	DoubleRtoI := map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
	res := 0
	for i := 0; i < length; i++ {
		if i+1 < length && DoubleRtoI[string(sb[i])+string(sb[i+1])] != 0 {
			res += DoubleRtoI[string(sb[i])+string(sb[i+1])]
			i = i + 1
		} else {
			res += RtoI[string(sb[i])]
		}

	}
	return res
}
func main() {
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
