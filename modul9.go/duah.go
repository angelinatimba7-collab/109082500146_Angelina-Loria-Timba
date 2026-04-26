package main

import (
	"fmt"
)

func main() {
	var n, x, idxHapus int
	fmt.Print("Masukkan jumlah elemen (N): ")
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Isi array:", arr)

	fmt.Print("Indeks ganjil: ")
	for i := 1; i < len(arr); i += 2 {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	for i := 0; i < len(arr); i++ {
		if i%x == 0 {
			fmt.Printf("%d ", arr[i])
		}
	}
	fmt.Println()

	sum := 0
	for _, v := range arr {
		sum += v
	}
	avg := float64(sum) / float64(len(arr))
	fmt.Printf("Rata-rata: %.2f\n", avg)

	fmt.Print("Hapus indeks: ")
	fmt.Scan(&idxHapus)
	arr = append(arr[:idxHapus], arr[idxHapus+1:]...)
	fmt.Println("Array setelah dihapus:", arr)
}
