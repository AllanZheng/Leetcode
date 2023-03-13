package main

import "fmt"

func convertTemperature(celsius float64) []float64 {
	return []float64{celsius + 273.15, celsius*1.8 + 32.00}
}

func main() {
	fmt.Println(convertTemperature(122.11234))
}
