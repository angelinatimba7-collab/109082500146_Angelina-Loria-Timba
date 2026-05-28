
# <h1 align="center">Laporan Praktikum Modul 14 - ... </h1>
<p align="center">[Angelina Loria Timba] - [109082500146]</p>

## Unguided 

### 1. Hercules, preman terkenal seantero ibukota, memiliki kerabat di banyak daerah. Tentunya Hercules sangat suka mengunjungi semua kerabatnya itu. Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program rumahkerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut membesar menggunakan algoritma selection sort. Masukan dimulai dengan sebuah integer n (0 < n < 1000), banyaknya daerah kerabat Hercules tinggal. Isi n baris berikutnya selalu dimulai dengan sebuah integer m (0 < m < 1000000) yang menyatakan banyaknya rumah kerabat di daerah tersebut, diikuti dengan rangkaian bilangan bulat positif, nomor rumah para kerabat. Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar di masing- masing daerah.
#### soal1.go

```go
package main

import "fmt"

const NMAX int = 1000

type arrRumah [1000000]int

func selectionSortAscending(A *arrRumah, n int) {
	var i, j, idxMin, temp int

	for i = 0; i < n-1; i++ {
		idxMin = i

		for j = i + 1; j < n; j++ {
			if A[j] < A[idxMin] {
				idxMin = j
			}
		}

		temp = A[i]
		A[i] = A[idxMin]
		A[idxMin] = temp
	}
}

func main() {
	var nDaerah, m, i, j int
	var rumah arrRumah

	fmt.Scan(&nDaerah)

	for i = 0; i < nDaerah; i++ {

		fmt.Scan(&m)

		for j = 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}

		selectionSortAscending(&rumah, m)

		for j = 0; j < m; j++ {
			fmt.Print(rumah[j])

			if j != m-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/angelinatimba7-collab/109082500146_Angelina-Loria-Timba/blob/main/modul14selectionsort.go/output/nomorsatu.png)
[Codingan ini dibuat untuk membantu Hercules menyusun nomor rumah kerabatnya agar lebih rapi dan mudah dikunjungi. Dengan algoritma Selection Sort, setiap nomor rumah diurutkan dari yang paling kecil hingga terbesar secara bertahap.]


## Unguided 

### 2. [ Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil. Format Masukan masih persis sama seperti sebelumnya. Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar untuk nomor ganjil, diikuti dengan terurut mengecil untuk nomor genap, di masing-masing daerah.]
#### soal2.go

```go
package main

import "fmt"

const NMAX int = 1000000

type arrInt [NMAX]int

func selectionSortAscending(A *arrInt, n int) {
	var i, j, idxMin, temp int

	for i = 0; i < n-1; i++ {
		idxMin = i

		for j = i + 1; j < n; j++ {
			if A[j] < A[idxMin] {
				idxMin = j
			}
		}

		temp = A[i]
		A[i] = A[idxMin]
		A[idxMin] = temp
	}
}

func selectionSortDescending(A *arrInt, n int) {
	var i, j, idxMax, temp int

	for i = 0; i < n-1; i++ {
		idxMax = i

		for j = i + 1; j < n; j++ {
			if A[j] > A[idxMax] {
				idxMax = j
			}
		}

		temp = A[i]
		A[i] = A[idxMax]
		A[idxMax] = temp
	}
}

func main() {
	var nDaerah, m int
	var i, j int

	var ganjil, genap arrInt
	var gCount, eCount int
	var x int

	fmt.Scan(&nDaerah)

	for i = 0; i < nDaerah; i++ {

		gCount = 0
		eCount = 0

		fmt.Scan(&m)

		for j = 0; j < m; j++ {

			fmt.Scan(&x)

			if x%2 == 1 {
				ganjil[gCount] = x
				gCount++
			} else {
				genap[eCount] = x
				eCount++
			}
		}

		selectionSortAscending(&ganjil, gCount)
		selectionSortDescending(&genap, eCount)

		for j = 0; j < gCount; j++ {
			fmt.Print(ganjil[j])

			if j != gCount-1 || eCount > 0 {
				fmt.Print(" ")
			}
		}

		for j = 0; j < eCount; j++ {
			fmt.Print(genap[j])

			if j != eCount-1 {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/angelinatimba7-collab/109082500146_Angelina-Loria-Timba/blob/main/modul2.go/output/soaldua.png)
[Saya membuat ini untuk mengatur nomor rumah berdasarkan sisi jalan agar Hercules tidak perlu sering menyeberang. Nomor ganjil akan ditampilkan lebih dulu secara menaik, kemudian dilanjutkan nomor genap secara menurun menggunakan proses Selection Sort.]

## Unguided 

### 3. [Kompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruan tinggi ternama. Dalam kompetisi tersebut, setiap tim berlomba untuk menyelesaikan sebanyak mungkin problem yang diberikan. Dari 13 problem yang diberikan, ada satu problem yang menarik. Problem tersebut mudah dipahami, hampir semua tim mencoba untuk menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemnya? "Median adalah nilai tengah dari suatu koleksi data yang sudah terurut. Jika jumlah data genap, maka nilai median adalah rerata dari kedua nilai tengahnya. Pada problem ini, semua data merupakan bilangan bulat positif, dan karenanya rerata nilai tengah dibulatkan ke bawah." Buatlah program median yang mencetak nilai median terhadap seluruh data yang sudah terbaca, jika data yang dibaca saat itu adalah 0. Masukan berbentuk rangkaian bilangan bulat. Masukan tidak akan berisi lebih dari 1000000 data, tidak termasuk bilangan 0. Data 0 merupakan tanda bahwa median harus dicetak, tidak termasuk data yang dicari mediannya. Data masukan diakhiri dengan bilangan bulat -5313. Keluaran adalah median yang diminta, satu data per baris.]
#### soal3.go

```go
package main

import "fmt"

const NMAX int = 1000000

type arrInt [NMAX]int

func insertionSort(A *arrInt, n int) {
	var i, j, temp int

	for i = 1; i < n; i++ {

		temp = A[i]
		j = i

		for j > 0 && temp < A[j-1] {
			A[j] = A[j-1]
			j--
		}

		A[j] = temp
	}
}

func hitungMedian(A arrInt, n int) int {

	if n%2 == 1 {
		return A[n/2]
	}

	return (A[(n/2)-1] + A[n/2]) / 2
}

func main() {

	var data arrInt
	var x, n int

	n = 0

	for {

		fmt.Scan(&x)

		if x == -5313 {
			break
		}

		if x == 0 {

			insertionSort(&data, n)

			fmt.Println(hitungMedian(data, n))

		} else {

			data[n] = x
			n++
		}
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/angelinatimba7-collab/109082500146_Angelina-Loria-Timba/blob/main/modul2.go/output/soaltiga.png)
[Program ini digunakan untuk menentukan nilai median dari kumpulan data yang terus bertambah. Setiap kali angka 0 dibaca, data akan diurutkan menggunakan Insertion Sort lalu dicari nilai tengahnya sesuai ketentuan soal.]
