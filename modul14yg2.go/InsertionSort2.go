package main

import "fmt"

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku struct {
	pustaka  [nMax]Buku
	nPustaka int
}

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	pustaka.nPustaka = n
	for i := 0; i < n; i++ {
		fmt.Scan(&pustaka.pustaka[i].id, &pustaka.pustaka[i].judul, &pustaka.pustaka[i].penulis,
			&pustaka.pustaka[i].penerbit, &pustaka.pustaka[i].eksemplar, &pustaka.pustaka[i].tahun,
			&pustaka.pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		return
	}

	idxFav := 0
	for i := 1; i < n; i++ {
		if pustaka.pustaka[i].rating > pustaka.pustaka[idxFav].rating {
			idxFav = i
		}
	}
	b := pustaka.pustaka[idxFav]
	fmt.Printf("Terfavorit: %s %s %s %d\n", b.judul, b.penulis, b.penerbit, b.tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {

	for i := 1; i < n; i++ {
		temp := pustaka.pustaka[i]
		j := i
		for j > 0 && temp.rating > pustaka.pustaka[j-1].rating {
			pustaka.pustaka[j] = pustaka.pustaka[j-1]
			j--
		}
		pustaka.pustaka[j] = temp
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	limit := 5
	if n < limit {
		limit = n
	}
	fmt.Println("5 Buku Teratas:")
	for i := 0; i < limit; i++ {
		fmt.Println(pustaka.pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {

	kiri := 0
	kanan := n - 1
	ketemu := false

	for kiri <= kanan && !ketemu {
		tengah := (kiri + kanan) / 2
		if pustaka.pustaka[tengah].rating == r {
			b := pustaka.pustaka[tengah]
			fmt.Printf("Buku Ditemukan: %s %s %s %d %d %d\n", b.judul, b.penulis, b.penerbit, b.tahun, b.eksemplar, b.rating)
			ketemu = true
		} else if r > pustaka.pustaka[tengah].rating {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if !ketemu {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var perpustakaan DaftarBuku
	var n, ratingCari int

	fmt.Scan(&n)
	DaftarkanBuku(&perpustakaan, n)

	if n > 0 {
		fmt.Scan(&ratingCari)
		CetakTerfavorit(perpustakaan, n)
		UrutBuku(&perpustakaan, n)
		Cetak5Terbaru(perpustakaan, n)
		CariBuku(perpustakaan, n, ratingCari)
	}
}
