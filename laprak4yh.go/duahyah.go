package main

import (
	"fmt"
)

func hitungSkor(waktu []int) (int, int) {
	soal := 0
	total := 0

	for _, w := range waktu {
		if w < 301 {
			soal++
			total += w
		}
	}
	return soal, total
}

func main() {
	var nama, pemenang string
	var bestSoal, bestWaktu int

	for {
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}

		waktu := make([]int, 8)
		for i := 0; i < 8; i++ {
			fmt.Scan(&waktu[i])
		}

		soal, total := hitungSkor(waktu)

		if soal > bestSoal || (soal == bestSoal && total < bestWaktu) {
			bestSoal = soal
			bestWaktu = total
			pemenang = nama
		}
	}

	fmt.Println(pemenang, bestSoal, bestWaktu)
}
