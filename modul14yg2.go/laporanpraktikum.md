
# <h1 align="center">Laporan Praktikum Modul 14(2) - ... </h1>
<p align="center">[Angelina Loria Timba] - [109082500146]</p>

## Unguided 

### 1. Buatlah sebuah program yang digunakan untuk membaca data integer seperti contoh yang diberikan di bawah ini, kemudian diurutkan (menggunakan metoda insertion sort), dan memeriksa apakah data yang terurut berjarak sama terhadap data sebelumnya. Masukan terdiri dari sekumpulan bilangan bulat yang diakhiri oleh bilangan negatif. Hanya bilangan non negatif saja yang disimpan ke dalam array. Keluaran terdiri dari dua baris. Baris pertama adalah isi dari array setelah dilakukan pengurutan, sedangkan baris kedua adalah status jarak setiap bilangan yang ada di dalam array. "Data berjarak x" atau "data berjarak tidak tetap".
#### soal1.go

```go
package main

import "fmt"

func main() {
	var arr []int
	var val int

	for {
		fmt.Scan(&val)
		if val < 0 {
			break
		}
		arr = append(arr, val)
	}

	n := len(arr)

	for i := 1; i < n; i++ {
		temp := arr[i]
		j := i
		for j > 0 && temp < arr[j-1] {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = temp
	}

	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(arr[i])
	}
	fmt.Println()

	if n <= 1 {
		fmt.Println("Data berjarak tidak tetap")
		return
	}

	jarak := arr[1] - arr[0]
	tetap := true
	for i := 1; i < n-1; i++ {
		if arr[i+1]-arr[i] != jarak {
			tetap = false
			break
		}
	}

	if tetap {
		fmt.Printf("Data berjarak %d\n", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/angelinatimba7-collab/109082500146_Angelina-Loria-Timba/blob/main/modul2.go/output/soalsatu.png)
[Kode ini membaca serangkaian angka hingga mendapati bilangan negatif, lalu menyisipkan hanya angka-angka positif ke dalam array untuk diurutkan dengan Insertion Sort. Setelah data urut, program akan melakukan iterasi untuk mengecek apakah selisih (jarak) antar elemen yang bersebelahan selalu konstan atau tidak.]


## Unguided 

### 2. [ Sebuah program perpustakaan digunakan untuk mengelola data buku di dalam suatu perpustakaan. Misalnya terdefinisi struct dan array seperti berikut ini: const nMax : integer = 7919 type Buku = < id, judul, penulis, penerbit : string eksemplar, tahun, rating : integer > type DaftarBuku = array [ 1..nMax] of Buku Pustaka : DaftarBuku nPustaka: integer Masukan terdiri dari beberapa baris. Baris pertama adalah bilangan bulat N yang menyatakan banyaknya data buku yang ada di dalam perpustakaan. N baris berikutnya, masing-masingnya adalah data buku sesuai dengan atribut atau field pada struct. Baris terakhir adalah bilangan bulat yang menyatakan rating buku yang akan dicari. Halaman 103 | Modul P raktikum Algoritma dan Pemrograman 2 Keluaran terdiri dari beberapa baris. Baris pertama adalah data buku terfavorit, baris kedua adalah lima judul buku dengan rating tertinggi, selanjutnya baris terakhir adalah data buku yang dicari sesuai rating yang diberikan pada masukan baris terakhir. Lengkapi subprogram-subprogram dibawah ini, sesuai dengan I.S. dan F.S yang diberikan. procedure DaftarkanBuku(in/out pustaka : DaftarBuku, n : integer) {I.S. sejumlah n data buku telah siap para piranti masukan F.S. n berisi sebuah nilai, dan pustaka berisi sejumlah n data buku} procedure CetakTerfavorit(in pustaka : DaftarBuku, in n : integer) {I.S. array pustaka berisi n buah data buku dan belum terurut F.S. Tampilan data buku (judul, penulis, penerbit, tahun)terfavorit, yaitu memiliki rating tertinggi} procedure UrutBuku( in/out pustaka : DaftarBuku, n : integer ) {I.S. Array pustaka berisi n data buku F.S. Array pustaka terurut menurun/mengecil terhadap rating. Catatan: Gunakan metoda Insertion sort} procedure Cetak5Terbaru( in pustaka : DaftarBuku, n integer ) {I.S. pustaka berisi n data buku yang sudah terurut menurut rating F.S. Laporan 5 judul buku dengan rating tertinggi Catatan: Isi pustaka mungkin saja kurang dari 5} procedure CariBuku(in pustaka : DaftarBuku, n : integer, r : integer ) {I.S. pustaka berisi n data buku yang sudah terurut menurut rating F.S. Laporan salah satu buku (judul, penulis, penerbit, tahun, eksemplar, rating) dengan rating yang diberikan. Jika tidak ada buku dengan rating yang ditanyakan, cukup tuliskan “Tidak ada buku dengan rating seperti itu”. Catatan: Gunakan pencarian biner/belah dua.}]
#### soal2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/angelinatimba7-collab/109082500146_Angelina-Loria-Timba/blob/main/modul2.go/output/soaldua.png)
[Program pengelolaan perpustakaan ini dipecah ke dalam beberapa prosedur, persis mengikuti struktur struct dan subprogram yang diminta soal. Bagian intinya terletak pada prosedur UrutBuku yang mengimplementasikan Insertion Sort secara mengecil (descending) berdasarkan rating, dan fungsi pencarian CariBuku yang menggunakan pola Binary Search untuk menemukan buku.]
