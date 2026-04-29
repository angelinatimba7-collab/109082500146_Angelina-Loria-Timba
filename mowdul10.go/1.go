package main

import "fmt"

func main() {
	var N int
	var berat [1000]float64
	var min, max float64

	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&berat[i])
		if i == 0 {
			min = berat[i]
			max = berat[i]
		}
		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}
	}
	fmt.Printf("%.2f %.2f\n", min, max)
}
