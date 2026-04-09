# <h1 align="center">Laporan Praktikum Modul 5 - ... </h1>
<p align="center">[Angelina Loria Timba] - [109082500146]</p>

## Unguided 

### 1. Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilai suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan Sn = Sn−1 + Sn−2 . Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci tersebut.
#### soal1.go

```go
package main

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah suku: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print(fib(i), " ")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/angelinatimba7-collab/109082500146_Angelina-Loria-Timba/blob/main/mod3smstdua.go/output/soalsatwu.png)
[Program ini dipakai buat nampilin deret Fibonacci pakai cara rekursif. Jadi fungsi fib bakal manggil dirinya sendiri buat dapetin nilai berikutnya berdasarkan dua angka sebelumnya. User masukin berapa banyak suku yang mau ditampilin, terus program bakal looping dari 0 sampai n dan nampilin hasilnya satu per satu. Ada kondisi berhenti di n = 0 dan n = 1 biar rekursinya ga jalan terus.]


## Unguided 

### 2. Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan menggunakan fungsi rekursif. N adalah masukan dari user.
#### soal2.go

```go
package main

import "fmt"

func bintang(n, i int) {
	if i > n {
		return
	}

	for j := 0; j < i; j++ {
		fmt.Print("*")
	}
	fmt.Println()

	bintang(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)

	bintang(n, 1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/angelinatimba7-collab/109082500146_Angelina-Loria-Timba/blob/main/mod3smstdua.go/output/dwuwah.png)
[Program ini buat nampilin pola bintang naik ke bawah (segitiga) pakai rekursi. Jadi awalnya mulai dari 1 bintang, terus tiap baris nambah satu sampai sesuai angka yang dimasukin user. Fungsi bintang bakal dipanggil terus dengan nilai i yang bertambah sampai lebih dari n. Jadi prosesnya kayak bertahap gitu, dari sedikit bintang sampai banyak.]

## Unguided 

### 3. Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N. Masukan terdiri dari sebuah bilangan bulat positif N. Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N (terurut dari 1 hingga N ya).
#### soal3.go

```go
package main

import "fmt"

func faktor(n, i int) {
	if i > n {
		return
	}

	if n%i == 0 {
		fmt.Print(i, " ")
	}

	faktor(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	fmt.Print("Faktor: ")
	faktor(n, 1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/angelinatimba7-collab/109082500146_Angelina-Loria-Timba/blob/main/mod3smstdua.go/output/tigwawh.png)
[Program ini dipakai buat nyari faktor dari suatu bilangan. Cara kerjanya ngecek dari angka 1 sampai n, terus kalau angka itu bisa ngebagi n tanpa sisa, berarti dia faktor dan bakal ditampilin. Semua dicek pakai rekursi, jadi fungsi bakal manggil dirinya sendiri sampai semua angka selesai dicek. Hasilnya nanti muncul faktor-faktor dari kecil ke besar.]
