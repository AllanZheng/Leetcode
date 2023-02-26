package main

func validPalindrome(s string) bool {
	replace := false
	length := len(s)
	sb := []byte(s)
	var half int
	if length%2 == 0 {
		half = length / 2
	} else {
		half = length/2 + 1
	}
	var sb1 []byte
	sb1 = []byte(s)
	front := false
	end := false
	for i := 0; i < half; i++ {
		if sb[i] != sb[length-1-i] && sb1[i] != sb1[length-1-i] {
			if replace == true {
				return false
			} else {
				if sb[i] != sb[length-2-i] && sb[i+1] != sb[length-1-i] {
					return false
				} else {

					sb = append(sb[0:i], sb[i+1:length]...)

					sb1 = append(sb1[0:length-1-i], sb1[length-i:length]...)
					//println(i, length-1-i)
					//println(string(s))
					//println(string(sb))
					//println(string(sb1))
					length = length - 1
					i = i - 1
					replace = true
				}

			}
		} else if replace == true {
			if sb[i] != sb[length-1-i] {
				front = true
			}
			if sb1[i] != sb1[length-1-i] {
				end = true
			}
			if front && end {
				return false
			}
		}
	}
	return true
}

func main() {
	s := "aec"
	print(validPalindrome(s))
}
