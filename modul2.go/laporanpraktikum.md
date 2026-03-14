
# <h1 align="center">Laporan Praktikum Modul 2 - ... </h1>
<p align="center">[Angelina Loria Timba] - [109082500146]</p>

## Unguided 

### 1. Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan
[Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan
masukan data yang sesuai sebanyak yang diminta program. Perhatikan keluaran yang
diperoleh. Coba terangkan apa sebenarnya yang dilakukan program tersebut?]
#### soal1.go

```go
package main
import "fmt"

func main() {

	var satu, dua, tiga string
	var temp string

	fmt.Print("Masukkan string pertama: ")
	fmt.Scanln(&satu)

	fmt.Print("Masukkan string kedua: ")
	fmt.Scanln(&dua)

	fmt.Print("Masukkan string ketiga: ")
	fmt.Scanln(&tiga)

	fmt.Println("Output awal =", satu, dua, tiga)

	temp = satu
	satu = dua
	dua = tiga
	tiga = temp

	fmt.Println("Output akhir =", satu, dua, tiga)

}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/angelinatimba7-collab/109082500146_Angelina_Loria_Timba/blob/main/output/soalsatu.png)
[Program ini menerima tiga string lalu menukar urutannya menggunakan variabel sementara temp. Awalnya urutan satu dua tiga, lalu setelah dipindahkan nilainya menjadi dua tiga satu, kemudian hasilnya ditampilkan.]


## Unguided 

### 2. Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan
[Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan
praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana
susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta
untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan
warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah ‘merah’, ‘kuning’,
‘hijau’, dan ‘ungu’ selama 5 kali percobaan berulang.
Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi
sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan
warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false
untuk urutan warna lainnya.]
#### soal2.go

```go
package main
import "fmt"

func main() {

	var warna1, warna2, warna3, warna4 string
	berhasil := true

	for i := 1; i <= 5; i++ {

		fmt.Print("Percobaan ", i, ": ")
		fmt.Scan(&warna1, &warna2, &warna3, &warna4)

		if warna1 != "merah" || warna2 != "kuning" || warna3 != "hijau" || warna4 != "ungu" {
			berhasil = false
		}

	}

	fmt.Println("BERHASIL:", berhasil)

}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/angelinatimba7-collab/109082500146_Angelina_Loria_Timba/blob/main/output/soalsatu.png)
[Codingan ini mengecek keberhasilan percobaan warna tabung reaksi. Input warna dimasukkan lima kali menggunakan for, lalu dicek apakah urutannya merah, kuning, hijau, ungu. Jika ada yang salah hasilnya false, jika semua benar true.]

## Unguided 

### 3. Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan
[PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka,
buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan
sebagai berikut!
Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam
gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500
gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500
gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang
kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.]
#### soal3.go

```go
package main
import "fmt"

func main() {

	var berat int
	var kg, gram int
	var biayaKg, biayaGram, total int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat)
	kg = berat / 1000
	gram = berat % 1000
	biayaKg = kg * 10000
	if kg > 10 {
		biayaGram = 0
	} else {

		if gram >= 500 {
			biayaGram = gram * 5
		} else {
			biayaGram = gram * 15
		}

	}
	total = biayaKg + biayaGram
	fmt.Println("Detail berat:", kg, "kg +", gram, "gr")
	fmt.Println("Detail biaya: Rp.", biayaKg, "+ Rp.", biayaGram)
	fmt.Println("Total biaya: Rp.", total)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/angelinatimba7-collab/109082500146_Angelina_Loria_Timba/blob/main/output/soalsatu.png)
[Programnya  menghitung biaya pengiriman berdasarkan berat parsel dalam gram. Berat diubah ke kilogram dan sisa gram, lalu biaya dihitung sesuai aturan dan ditampilkan totalnya.]

