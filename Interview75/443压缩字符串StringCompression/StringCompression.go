package main

import (
	"fmt"
)

func compress1(chars []byte) int {
	result := 1
	count := 1
	if len(chars) == 1 {
		return 1
	}
	for i := 1; i < len(chars); i++ {
		if chars[i] == chars[i-1] {
			count = count + 1
		} else {
			if count == 1 {
				result = result + 1
			} else {
				result = result + 2
			}

			for count/10 >= 1 {

				result++
				count = count / 10

			}
			count = 1
		}
	}
	result++
	for count/10 >= 1 {

		result++
		count = count / 10

	}

	return result
}
func compress(chars []byte) int {

	if len(chars) == 1 {
		return 1
	}
	count := 1
	move := []byte{}
	start := 1
	end := 1
	for i := 1; i < len(chars); i++ {
		if chars[i] == chars[i-1] {
			start = i
			count++
		} else {
			end = i - 1
			//back := chars[:end]
			for count > 1 {
				move = append(move)
			}
			chars = append(chars[start:], chars[:end]...)
		}
		count = 1
		i = start + 1
	}
	return len(chars)
}

func compress3(chars []byte) int {

	write, left := 0, 0
	for read, ch := range chars {
		if read == len(chars)-1 || ch != chars[read+1] {
			chars[write] = ch
			write++
			num := read - left + 1
			if num > 1 {
				anchor := write
				for ; num > 0; num /= 10 {
					chars[write] = '0' + byte(num%10)
					write++
				}
				s := chars[anchor:write]
				for i, n := 0, len(s); i < n/2; i++ {
					s[i], s[n-1-i] = s[n-1-i], s[i]
				}
			}
			left = read + 1
		}

	}
	return write
}

func main() {
	fmt.Println(compress3([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}))
	fmt.Println(compress3([]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}))
	fmt.Println(compress3([]byte{'a'}))
}
