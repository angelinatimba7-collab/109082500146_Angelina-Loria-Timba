# <h1 align="center">Laporan Praktikum Modul 9 - ... </h1>
<p align="center">[Angelina Loria Timba] - [109082500146]</p>

## Unguided 

### 1. Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radiusnya. Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat. Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".
#### soal1.go

```go
package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}

type Lingkaran struct {
	pusat  Titik
	radius int
}

func jarak(p, q Titik) float64 {
	return math.Sqrt(math.SmallestNonzeroFloat64 + float64((p.x-q.x)*(p.x-q.x)+(p.y-q.y)*(p.y-q.y)))
}

func didalam(c Lingkaran, p Titik) bool {
	return jarak(c.pusat, p) <= float64(c.radius)
}

func main() {
	var l1, l2 Lingkaran
	var p Titik

	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.radius)
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.radius)
	fmt.Scan(&p.x, &p.y)

	inL1 := didalam(l1, p)
	inL2 := didalam(l2, p)

	if inL1 && inL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/angelinatimba7-collab/109082500146_Angelina-Loria-Timba/blob/main/mod3smstdua.go/output/soalsatwu.png)
[Program ini menentukan posisi sebuah titik terhadap dua lingkaran menggunakan tipe bentukan (struct) titik dan lingkaran. Penentuan posisi dilakukan dengan menghitung jarak titik ke pusat lingkaran menggunakan rumus Euclidean.]


## Unguided 

### 2. Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program yang digunakan untuk mengisi array tersebut sebanyak N elemen nilai. Asumsikan array memiliki kapasitas penyimpanan data sejumlah elemen tertentu.
#### soal2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/angelinatimba7-collab/109082500146_Angelina-Loria-Timba/blob/main/mod3smstdua.go/output/dwuwah.png)
[ Program ini mengelola sekumpulan bilangan bulat dalam array untuk menampilkan data berdasarkan indeks ganjil, genap, atau kelipatan tertentu. Selain itu, program juga menghitung nilai statistik seperti rata-rata, standar deviasi, dan frekuensi kemunculan angka.]

## Unguided 

### 3. Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga. Pertama-tama program meminta masukan nama-nama klub yang bertanding, kemudian program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan dalam array adalah nama-nama klub yang menang saja. Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir program, tampilkan daftar klub yang memenangkan pertandingan.
#### soal3.go

```go
package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var pemenang []string

	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	i := 1
	for {
		fmt.Printf("Pertandingan %d: ", i)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			pemenang = append(pemenang, klubA)
			fmt.Printf("Hasil %d: %s\n", i, klubA)
		} else if skorB > skorA {
			pemenang = append(pemenang, klubB)
			fmt.Printf("Hasil %d: %s\n", i, klubB)
		} else {
			fmt.Printf("Hasil %d: Draw\n", i)
		}
		i++
	}

	fmt.Println("Daftar klub yang menang:", pemenang)
	fmt.Println("Pertandingan selesai")
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/angelinatimba7-collab/109082500146_Angelina-Loria-Timba/blob/main/mod3smstdua.go/output/tigwawh.png)
[Program ini merekap hasil pertandingan bola dengan menyimpan nama klub yang menang ke dalam sebuah array. Proses input data skor akan terus berjalan dan baru akan berhenti jika pengguna memasukkan angka negatif.]

## Unguided 

### 4. Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa apakah membentuk palindrom.
#### soal3.go

```go
package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var input rune
	*n = 0
	for *n < NMAX {
		fmt.Scanf("%c", &input)
		if input == '.' || input == '\n' {
			break
		}
		if input != ' ' {
			t[*n] = input
			*n++
		}
	}
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	asli := t
	balikanArray(&t, n)
	for i := 0; i < n; i++ {
		if asli[i] != t[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Teks: ")
	isiArray(&tab, &m)

	isPal := palindrom(tab, m)

	balikanArray(&tab, m)
	fmt.Print("Reverse teks: ")
	for i := 0; i < m; i++ {
		fmt.Printf("%c", tab[i])
	}

	fmt.Printf("\nPalindrom? %t\n", isPal)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/angelinatimba7-collab/109082500146_Angelina-Loria-Timba/blob/main/mod3smstdua.go/output/tigwawh.png)
[Program ini memproses sekumpulan karakter dalam array untuk membalikkan urutan isinya dan memeriksa apakah susunan tersebut membentuk palindrom. Proses input karakter akan berhenti jika pengguna memasukkan karakter titik atau kapasitas array sudah penuh.]