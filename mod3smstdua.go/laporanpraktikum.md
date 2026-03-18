# <h1 align="center">Laporan Praktikum Modul 2 - ... </h1>
<p align="center">[Angelina Loria Timba] - [109082500146]</p>

## Unguided 

### 1. yang sesuai sebanyak yang diminta program. Perhatikan keluaran yang
Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika
diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng
untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian
membantu Jonas? (tidak tentunya ya :p). Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi,
dengan syarat a ≥ c dan b ≥ d.
Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a
terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d.
#### soal1.go

```go
package main
import "fmt"

func fact(n int) int {
    res := 1
    for i := 1; i <= n; i++ {
        res *= i
    }
    return res
}

func perm(n, r int) int {
    return fact(n) / fact(n-r)
}

func comb(n, r int) int {
    return fact(n) / (fact(r) * fact(n-r))
}

func main() {
    var a, b, c, d int
    fmt.Scan(&a, &b, &c, &d)

    fmt.Println(perm(a,c), comb(a,c))
    fmt.Println(perm(b,d), comb(b,d))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/angelinatimba7-collab/109082500146_Angelina-Loria-Timba/blob/main/mod3smstdua.go/output/soalsatwu.png)
[Bikin fungsi faktorial dulu, terus pakai buat hitung permutasi (perm) sama kombinasi (comb). Di main tinggal masukin angka terus print hasilnya. Simple, langsung kelihatan berapa banyak cara atau kombinasi yang mungkin.]


## Unguided 

### 2. Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan
[Diberikan tiga buah fungsi matematika yaitu f (x) = x
2
, g (x) = x − 2 dan h (x) = x +
1. Fungsi komposisi (fogoh)(x) artinya adalah f(g(h(x))). Tuliskan f(x), g(x) dan h(x)
dalam bentuk function.
Masukan terdiri dari sebuah bilangan bulat a, b dan c yang dipisahkan oleh spasi.
Keluaran terdiri dari tiga baris. Baris pertama adalah (fogoh)(a), baris kedua (gohof)(b),
dan baris ketiga adalah (hofog)(c)!]
#### soal2.go

```go
package main
import "fmt"

func f(x int) int { return x*x }
func g(x int) int { return x-2 }
func h(x int) int { return x+1 }

func fogoh(x int) int { return f(g(h(x))) }
func gohof(x int) int { return g(h(f(x))) }
func hofog(x int) int { return h(f(g(x))) }

func main() {
    var a, b, c int
    fmt.Scan(&a, &b, &c)
    fmt.Println(fogoh(a))
    fmt.Println(gohof(b))
    fmt.Println(hofog(c))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/angelinatimba7-collab/109082500146_Angelina-Loria-Timba/blob/main/mod3smstdua.go/output/dwuwah.png)
[Bikin fungsi dasar f, g, h, terus bikin versi komposisinya sesuai soal (f∘g∘h, dll). Input dimasukin, fungsi dipanggil, hasilnya keluar langsung. Gampang dimengerti karena tinggal ngikutin urutan fungsi.]

## Unguided 

### 3. Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan
[[Lingkaran] Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius
r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y)
berdasarkan dua lingkaran tersebut.
Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat
dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik
sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan
bilangan bulat.
Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik
di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2"..]
#### soal3.go

```go
package main
import (
	"fmt"
	"math"
)
func jarak(x1, y1, x2, y2 float64) float64 {
	return math.Hypot(x1-x2, y1-y2)
}
func diDalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}
func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	in1 := diDalam(cx1, cy1, r1, x, y)
	in2 := diDalam(cx2, cy2, r2, x, y)

	switch {
	case in1 && in2:
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	case in1:
		fmt.Println("Titik di dalam lingkaran 1")
	case in2:
		fmt.Println("Titik di dalam lingkaran 2")
	default:
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/angelinatimba7-collab/109082500146_Angelina_Loria_Timba/blob/main/mod3smstdua/output/tigwawh.png)
[Bikin fungsi buat hitung jarak titik dan ngecek apakah titik itu ada di dalam lingkaran. Di main scan data lingkaran + titik, terus tinggal switch kondisi biar tahu titik ada di lingkaran 1, 2, keduanya, atau di luar.]
