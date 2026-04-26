# <h1 align="center">Laporan Praktikum Modul 4 - ... </h1>
<p align="center">[Angelina Loria Timba] - [109082500146]</p>

## Unguided 

### 1. yang sesuai sebanyak yang diminta program. Perhatikan keluaran yang Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p) Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a ≥ c dan b ≥ d. Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d.
#### soal1.go

```go
package main

import "fmt"

func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)

	fmt.Println(permutation(a, c), combination(a, c))
	fmt.Println(permutation(b, d), combination(b, d))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/angelinatimba7-collab/109082500146_Angelina-Loria-Timba/blob/main/laprak4yh.go/output/satu.png)
[Program ini intinya buat ngitung dua hal dari angka yang kita masukin, yaitu permutasi dan kombinasi. Dia punya fungsi buat faktorial dulu (perkalian dari 1 sampai n), karena dua rumus itu butuh faktorial. Setelah itu, program tinggal masukin angka ke rumus permutasi dan kombinasi, lalu hasilnya ditampilkan dalam dua baris sesuai pasangan angka yang diminta..]


## Unguided 

### 2. Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan soal paling banyak dalam waktu paling singkat adalah pemenangnya. Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur. prosedure hitungSkor(in/out soal, skor : integer) Setiap baris masukan dimulai dengan satu string nama peserta tersebut diikuti dengan adalah 8 integer yang menyatakan berapa lama (dalam menit) peserta tersebut menyelesaikan soal. Jika tidak berhasil atau tidak mengirimkan jawaban maka otomatis dianggap menyelesaikan dalam waktu 5 jam 1 menit (301 menit). Satu baris keluaran berisi nama pemenang, jumlah soal yang diselesaikan, dan nilai yang diperoleh. Nilai adalah total waktu yang dibutuhkan untuk menyelesaikan soal yang berhasil diselesaikan.
#### soal2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/angelinatimba7-collab/109082500146_Angelina-Loria-Timba/blob/main/laprak4yh.go/output/dua.png)
[2. Program ini kayak ngecek siapa yang menang lomba coding. Setiap orang punya waktu ngerjain 8 soal, dan kalau waktunya lebih dari 300 menit berarti dianggap gagal. Program bakal hitung berapa soal yang berhasil dikerjain dan total waktunya, lalu dibandingin sama peserta lain. Yang ngerjain soal paling banyak bakal menang, tapi kalau jumlahnya sama, yang waktunya paling cepat yang menang.]
