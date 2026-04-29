package main

import "fmt"

func main() {
	var x, y int
	var ikan [1000]float64

	fmt.Scan(&x, &y)
	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	var totalSeluruhWadah float64
	jumlahWadah := 0

	for i := 0; i < x; i += y {
		var totalWadah float64
		for j := i; j < i+y && j < x; j++ {
			totalWadah += ikan[j]
		}
		fmt.Printf("%.2f ", totalWadah)
		totalSeluruhWadah += totalWadah
		jumlahWadah++
	}

	fmt.Printf("\n%.2f\n", totalSeluruhWadah/float64(jumlahWadah))
}
