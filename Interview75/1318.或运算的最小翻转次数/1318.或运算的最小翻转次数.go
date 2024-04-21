package main

func minFlips(a int, b int, c int) int {
	res := 0

	for c > 0 {
		if c%2 == 0 {
			if b%2 != 0 {
				res++
			}
			if a%2 != 0 {
				res++
			}
		} else {
			if b%2 == 0 && a%2 == 0 {
				res++
			}

		}
		c = c / 2
		b = b / 2
		a = a / 2
	}
	for a > 0 {
		if a%2 == 1 {
			res++
		}
		a = a / 2
	}
	for b > 0 {
		if b%2 == 1 {
			res++
		}
		b = b / 2
	}
	return res
}
