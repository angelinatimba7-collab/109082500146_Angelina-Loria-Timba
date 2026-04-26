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
