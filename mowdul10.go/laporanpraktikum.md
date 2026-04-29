# <h1 align="center">Laporan Praktikum Modul 10 - ... </h1>
<p align="center">[Angelina Loria Timba] - [109082500146]</p>

## Unguided 

### 1. Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual. Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual. Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar.

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/angelinatimba7-collab/109082500146_Angelina-Loria-Timba/blob/main/mowdul10.go/output/1.png)
[Program ini menginisialisasi nilai minimum dan maksimum menggunakan elemen pertama yang dimasukkan oleh pengguna. Melalui sebuah perulangan, setiap input berat dibandingkan secara berurutan untuk memperbarui nilai terkecil dan terbesar.]


## Unguided 

### 2. Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual. Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil yang menyatakan banyaknya ikan yang akan dijual. Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang dimasukan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.
#### soal2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/angelinatimba7-collab/109082500146_Angelina-Loria-Timba/blob/main/mod3smstdua.go/output/dwuwah.png)
[Program menggunakan perulangan bertingkat untuk menjumlahkan berat ikan ke dalam kelompok-kelompok sesuai kapasitas wadah y. Hasil penjumlahan setiap wadah dicetak di baris pertama, sementara baris kedua menampilkan hasil bagi antara total seluruh berat dengan jumlah wadah yang digunakan.]

## Unguided 

### 3. Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.
#### soal3.go

```go
package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	var total float64
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var data arrBalita
	var min, max float64

	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}

	hitungMinMax(data, n, &min, &max)

	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rerata(data, n))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/angelinatimba7-collab/109082500146_Angelina-Loria-Timba/blob/main/mod3smstdua.go/output/tigwawh.png)
[Program ini mengimplementasikan prosedur hitungMinMax dengan parameter pointer untuk mengembalikan dua nilai sekaligus dan fungsi rerata untuk menghitung nilai rata-rata. Struktur kode ini memisahkan logika pengolahan data dari fungsi utama agar lebih terorganisir.]
